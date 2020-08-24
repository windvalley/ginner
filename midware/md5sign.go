package midware

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"use-gin/auth"
	"use-gin/config"
	"use-gin/errcode"
	"use-gin/handler"
)

func Md5sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		debugMsg, err := verifySign(c)
		if err != nil {
			err1 := errcode.New(errcode.APISignError, nil)
			err1.Add(err)
			handler.SendResponse(c, err1, nil)
			c.Abort()
			return
		}

		if debugMsg != nil {
			handler.SendResponse(c, nil, debugMsg)
			c.Abort()
			return
		}

		c.Next()
	}
}

type PublicParamsDebug struct {
	KeyID string `form:"KeyID" binding:"required"`
}

type PublicParams struct {
	PublicParamsDebug
	Timestamp string `form:"Timestamp" binding:"required"`
	Nonce     string `form:"Nonce" binding:"required"`
	Signature string `form:"Signature" binding:"required"`
}

func verifySign(c *gin.Context) (map[string]string, error) {
	debug := c.Query("debug")

	if err := c.Request.ParseForm(); err != nil {
		err1 := errcode.New(errcode.InternalServerError, err)
		handler.SendResponse(c, err1, nil)
		c.Abort()
		return nil, nil
	}
	allParamsMap := c.Request.Form

	var paramsDebug PublicParamsDebug
	var params PublicParams

	keyID := ""
	if debug == "1" {
		if err := c.ShouldBind(&paramsDebug); err != nil {
			err1 := errcode.New(errcode.ValidationError, err)
			err1.Add(err)
			handler.SendResponse(c, err1, nil)

			c.Abort()
			return nil, nil
		}
		keyID = paramsDebug.KeyID
	} else {
		if err := c.Bind(&params); err != nil {
			err1 := errcode.New(errcode.ValidationError, err)
			err1.Add(err)
			handler.SendResponse(c, err1, nil)

			c.Abort()
			return nil, nil
		}
		keyID = params.KeyID
	}

	userInfo, ok := auth.UserInfos[keyID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("KeyID '%s' not found", keyID))
	}

	keySecret := userInfo.MD5

	curTimestamp := time.Now().Unix()
	if debug == "1" {
		if config.Conf().Runmode == "debug" {
			rand.Seed(curTimestamp)
			nonce := strconv.Itoa(rand.Intn(100000))
			timestamp := strconv.FormatInt(curTimestamp, 10)
			allParamsMap.Set("Timestamp", timestamp)
			allParamsMap.Set("Nonce", nonce)
			strForSign := createStrForSign(c, allParamsMap)

			res := map[string]string{
				"Timestamp": timestamp,
				"Nonce":     nonce,
				"Signature": generateSign(strForSign, keySecret),
			}
			return res, nil
		}

		return nil, errors.New("debug forbidden in release runmode")
	}

	timestampReq, err := strconv.ParseInt(params.Timestamp, 10, 64)
	if err != nil {
		return nil, err
	}

	apisignLifetime := config.Conf().Auth.APISignLifetime

	if timestampReq > curTimestamp ||
		curTimestamp-timestampReq >= apisignLifetime {
		return nil, errors.New("Signature expired")
	}

	strForSign := createStrForSign(c, allParamsMap)
	if params.Signature == "" ||
		params.Signature != generateSign(strForSign, keySecret) {
		return nil, errors.New("Signature invalid")
	}

	return nil, nil
}

func createStrForSign(c *gin.Context, reqParamsMap url.Values) string {
	var keys []string
	for k := range reqParamsMap {
		if k != "Signature" && k != "debug" {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)

	params := ""
	for i := 0; i < len(keys); i++ {
		params = params + fmt.Sprintf("%v=%v", keys[i], reqParamsMap.Get(keys[i]))
	}

	reqMethod := c.Request.Method
	reqHost := strings.Split(c.Request.Host, ":")[0]
	reqPath := c.FullPath()

	return reqMethod + reqHost + reqPath + params
}

func generateSign(strForSign, keySecret string) string {
	signature := auth.Md5sum(keySecret + strForSign + keySecret)
	return signature
}

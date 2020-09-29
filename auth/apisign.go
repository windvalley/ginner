package auth

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

	"ginner/api"
	"ginner/config"
	"ginner/errcode"
)

type publicParamsDebug struct {
	KeyID string `form:"KeyID" binding:"required"`
}

type publicParams struct {
	publicParamsDebug
	Timestamp time.Time `form:"Timestamp" binding:"required" time_format:"unix"`
	Nonce     int       `form:"Nonce" binding:"required"`
	Signature string    `form:"Signature" binding:"required"`
}

// VerifySign verify if the signature of client is valid
func VerifySign(c *gin.Context, signType string) (map[string]string, error) {
	debug := c.Query("debug")

	if err := c.Request.ParseForm(); err != nil {
		err1 := errcode.New(errcode.InternalServerError, err)
		api.SendResponse(c, err1, nil)
		c.Abort()
		return nil, nil
	}
	allParamsMap := c.Request.Form

	var paramsDebug publicParamsDebug
	var params publicParams

	keyID := ""
	var err error
	if debug == "1" {
		err = c.ShouldBind(&paramsDebug)
		keyID = paramsDebug.KeyID
	} else {
		err = c.ShouldBind(&params)
		keyID = params.KeyID
	}
	if err != nil {
		err1 := errcode.New(errcode.ValidationError, nil)
		err1.Add(err)
		api.SendResponse(c, err1, nil)

		c.Abort()
		return nil, nil
	}

	userInfo, ok := userInfos[keyID]
	if !ok {
		return nil, fmt.Errorf("KeyID '%s' not found", keyID)
	}

	keySecret, err := getKeySecret(keyID, signType, userInfo)
	if err != nil {
		return nil, err
	}

	curTimestamp := time.Now().Unix()
	if debug == "1" {
		if config.Conf().Runmode == "debug" {
			rand.Seed(curTimestamp)
			nonce := strconv.Itoa(rand.Intn(100000))
			timestamp := strconv.FormatInt(curTimestamp, 10)
			allParamsMap.Set("Timestamp", timestamp)
			allParamsMap.Set("Nonce", nonce)
			strForSign := createStrForSign(c, allParamsMap)

			signature, err := generateSign(strForSign, keySecret, signType)
			if err != nil {
				return nil, err
			}

			res := map[string]string{
				"Timestamp": timestamp,
				"Nonce":     nonce,
				"Signature": signature,
			}
			return res, nil
		}

		return nil, errors.New("debug forbidden in release runmode")
	}

	apisignLifetime := config.Conf().Auth.APISignLifetime

	if params.Timestamp.Unix() > curTimestamp ||
		curTimestamp-params.Timestamp.Unix() >= apisignLifetime {
		return nil, errors.New("Signature expired")
	}

	strForSign := createStrForSign(c, allParamsMap)

	err = verifySiganature(strForSign, signType, keySecret, params.Signature, userInfo)

	return nil, err
}

func verifySiganature(strForSign, signType, keySecret, requestSignature string,
	userInfo userInfo) error {
	signInvalidError := errors.New("Signature invalid")

	switch signType {
	case "aes":
		srcStr, err := AESDecrypt(requestSignature, keySecret)
		if err != nil {
			return err
		}

		if srcStr != strForSign {
			return signInvalidError
		}
	case "rsa":
		srcStr, err := DecryptByPrivate(requestSignature, userInfo.RSA.Private)
		if err != nil {
			return err
		}

		if srcStr != strForSign {
			return signInvalidError
		}
	case "md5", "hmac_md5", "hmac_sha1", "hmac_sha256":
		signature, err := generateSign(strForSign, keySecret, signType)
		if err != nil {
			return err
		}

		if requestSignature != signature {
			return signInvalidError
		}
	default:
		return errors.New("Signature encrypt type invalid")
	}

	return nil
}

func getKeySecret(keyID, signType string, userInfo userInfo) (string, error) {
	keySecret := ""
	switch signType {
	case "md5":
		keySecret = userInfo.MD5
	case "aes":
		keySecret = userInfo.AES
	case "rsa":
		keySecret = userInfo.RSA.Public
	case "hmac_md5", "hmac_sha1", "hmac_sha256":
		keySecret = userInfo.Hmac
	}

	return keySecret, nil
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

func generateSign(strForSign, keySecret, signType string) (string, error) {
	signature := ""
	var err error
	switch signType {
	case "md5":
		signature = Md5sum(keySecret + strForSign + keySecret)
	case "aes":
		signature, err = AESEncrypt(strForSign, keySecret)
		if err != nil {
			return "", err
		}
	case "rsa":
		signature, err = EncryptByPublic(strForSign, keySecret)
		if err != nil {
			return "", err
		}
	case "hmac_md5", "hmac_sha1", "hmac_sha256":
		signature, err = Hmac(signType, strForSign, keySecret)
		if err != nil {
			return "", err
		}

	}

	return signature, nil
}

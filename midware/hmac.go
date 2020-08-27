package midware

import (
	"github.com/gin-gonic/gin"

	"use-gin/auth"
	"use-gin/errcode"
	"use-gin/handler"
)

func HmacMd5Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		debugMsg, err := auth.VerifySign(c, "hmac_md5")
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

func HmacSha1Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		debugMsg, err := auth.VerifySign(c, "hmac_sha1")
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

func HmacSha256Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		debugMsg, err := auth.VerifySign(c, "hmac_sha256")
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

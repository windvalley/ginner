package midware

import (
	"github.com/gin-gonic/gin"

	"use-gin/auth"
	"use-gin/errcode"
	"use-gin/handler"
)

func AESSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		debugMsg, err := auth.VerifySign(c, "aes")
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

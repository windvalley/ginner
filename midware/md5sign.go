package midware

import (
	"github.com/gin-gonic/gin"

	"use-gin/auth"
	"use-gin/errcode"
	"use-gin/handler"
)

// Md5Sign MD5 combined encryption signature
func Md5Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		debugMsg, err := auth.VerifySign(c, "md5")
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

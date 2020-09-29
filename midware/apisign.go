package midware

import (
	"github.com/gin-gonic/gin"

	"ginner/api"
	"ginner/auth"
	"ginner/errcode"
)

const (
	// SignTypeHmacMd5 hmac_md5
	SignTypeHmacMd5 = "hmac_md5"
	// SignTypeHmacSHA1 hmac_sha1
	SignTypeHmacSHA1 = "hmac_sha1"
	// SignTypeHmacSHA256 hmac_sha256
	SignTypeHmacSHA256 = "hmac_sha256"
	// SignTypeMd5 md5
	SignTypeMd5 = "md5"
	// SignTypeAES aes
	SignTypeAES = "aes"
	// SignTypeRSA rsa
	SignTypeRSA = "rsa"
)

// APISign API signature verify
func APISign(signType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		debugMsg, err := auth.VerifySign(c, signType)
		if err != nil {
			err1 := errcode.New(errcode.APISignError, nil)
			err1.Add(err)
			api.SendResponse(c, err1, nil)
			c.Abort()
			return
		}

		if debugMsg != nil {
			api.SendResponse(c, nil, debugMsg)
			c.Abort()
			return
		}

		c.Next()
	}
}

package midware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"use-gin/auth"
	"use-gin/errcode"
	"use-gin/handler"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.Query("jwt")

		if jwtToken == "" {
			err1 := errcode.New(errcode.TokenInvalidError, nil)
			err1.Add("no JWT(Json Web Token) found.")
			handler.SendResponse(c, err1, nil)

			c.Abort()
			return
		} else {
			if _, err := auth.ParseJWT(jwtToken); err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err1 := errcode.New(errcode.TokenInvalidError, err)
					err1.Add("JWT expired.")
					handler.SendResponse(c, err1, nil)
				case jwt.ValidationErrorSignatureInvalid:
					err1 := errcode.New(errcode.TokenInvalidError, err)
					err1.Add("JWT signature validation failed.")
					handler.SendResponse(c, err1, nil)
				default:
					err1 := errcode.New(errcode.TokenInvalidError, err)
					err1.Add("JWT validation failed.")
					handler.SendResponse(c, err1, nil)
				}

				c.Abort()
				return
			}
		}

		c.Next()
	}
}

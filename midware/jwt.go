package midware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"use-gin/auth"
	"use-gin/config"
	"use-gin/errcode"
	"use-gin/handler"
)

func JWT() gin.HandlerFunc {
	jwtSecret := config.Conf().Auth.JWTSecret

	return func(c *gin.Context) {
		jwtToken := c.Query("jwt")

		if jwtToken == "" {
			jwtToken = c.Query("token")
		}

		if jwtToken == "" {
			authHeader := c.GetHeader("Authorization")
			if authHeader != "" {
				values := strings.Split(authHeader, " ")
				if len(values) != 2 || values[0] != "Bearer" {
					err1 := errcode.New(errcode.TokenInvalidError, nil)
					err1.Add("Request header 'Authorization' format invalid.")
					handler.SendResponse(c, err1, nil)

					c.Abort()
					return
				}
				jwtToken = values[1]
			}
		}

		if jwtToken == "" {
			err1 := errcode.New(errcode.TokenInvalidError, nil)
			err1.Add("no JWT(Json Web Token) found.")
			handler.SendResponse(c, err1, nil)

			c.Abort()
			return
		} else {
			claims, err := auth.ParseJWT(jwtToken, jwtSecret)
			if err != nil {
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

			// NOTE: handler function can be get username by c.Get("username")
			c.Set("username", claims.Issuer)
		}

		c.Next()
	}
}

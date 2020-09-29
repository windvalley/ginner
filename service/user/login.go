package user

import (
	"github.com/jinzhu/gorm"

	"ginner/auth"
	"ginner/config"
	"ginner/errcode"
	"ginner/model"
)

// GetJWT get a jwt token if login success
func GetJWT(username, password string) (string, error) {
	d, err := model.GetUser(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errcode.UserNotFoundError
		}

		return "", errcode.New(errcode.InternalServerError, err)
	}

	if err := d.CheckPassword(password); err != nil {
		return "", errcode.New(errcode.PasswordIncorrectError, err)
	}

	secret := config.Conf().Auth.JWTSecret
	jwtLifetime := config.Conf().Auth.JWTLifetime
	jwt, err := auth.GenerateJWT(username, secret, jwtLifetime)
	if err != nil {
		return "", errcode.New(errcode.InternalServerError, err)
	}

	return jwt, nil
}

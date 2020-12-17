package service

import (
	"github.com/jinzhu/gorm"

	"ginner/auth"
	"ginner/config"
	"ginner/errcode"
	"ginner/model"
)

// CreateUser create user by username and password
func CreateUser(username, password string) error {
	u := &model.User{
		Username: username,
		Password: password,
	}

	if err := u.EncryptPassword(); err != nil {
		return err
	}

	err := u.Create()
	return err
}

// GetUser get user info by username
func GetUser(username string) (*model.User, error) {
	user, err := model.GetUser(username)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return user, err
}

// GetUserJWT get a jwt token if login success
func GetUserJWT(username, password string) (string, error) {
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

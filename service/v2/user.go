package service

import (
	"ginner/model"
)

// CreateUser create user by username and password
func CreateUser(username, password string) error {

	// omit logic

	return nil
}

// GetUser get user info by username
func GetUser(username string) (*model.User, error) {

	// omit logic

	return nil, nil
}

// GetUserJWT get a jwt token if login success
func GetUserJWT(username, password string) (string, error) {

	// omit logic

	return "", nil
}

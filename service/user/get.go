package user

import (
	"github.com/jinzhu/gorm"

	"use-gin/model"
)

// Get get user info by username
func Get(username string) (*model.User, error) {
	user, err := model.GetUser(username)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return user, err
}

// Package api The purpose of this module is just to demonstrate version control of the api
package api

import (
	"github.com/gin-gonic/gin"

	"ginner/service/v2"
)

// CreateUser user register
func CreateUser(c *gin.Context) {
	if err := service.CreateUser("username", "password"); err != nil {
		// omit
	}

	// omit
}

// Login user login
func Login(c *gin.Context) {
	_, err := service.GetUserJWT("username", "password")
	if err != nil {
		// omit
	}
	// omit
}

// GetUser get user by path params "username"
func GetUser(c *gin.Context) {
	_, err := service.GetUser("username")
	if err != nil {
		// omit
	}

	// omit
}

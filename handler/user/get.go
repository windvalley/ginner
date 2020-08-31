package user

import (
	"github.com/gin-gonic/gin"

	"use-gin/errcode"
	"use-gin/handler"
	"use-gin/model/rdb"
)

// GetUser get user by path params "username"
func GetUser(c *gin.Context) {
	username := c.Param("username")

	user, err := rdb.GetUser(username)
	if err != nil && err.Error() == "record not found" {
		handler.SendResponse(c, errcode.RecordNotFoundError, nil)
		return
	}

	if err != nil {
		err1 := errcode.New(errcode.DBError, err)
		err1.Add("get user failed.")
		handler.SendResponse(c, err1, nil)
		return
	}

	handler.SendResponse(c, nil, user)
}

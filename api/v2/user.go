package api

import (
	"github.com/gin-gonic/gin"

	"ginner/api"
	"ginner/errcode"
	"ginner/service/v2"
)

// CreateUser user register
func CreateUser(c *gin.Context) {
	var r userInfo
	if err := c.Bind(&r); err != nil {
		err1 := errcode.New(errcode.ValidationError, err)
		err1.Add(err)
		api.SendResponse(c, err1, nil)
		return
	}

	if err := service.CreateUser(r.Username, r.Password); err != nil {
		err1 := errcode.New(errcode.CustomInternalServerError, err)
		err1.Add("create user failed.")
		api.SendResponse(c, err1, nil)
		return
	}

	api.SendResponse(c, errcode.Created, &userCreateResp{r.Username})
}

// Login user login
func Login(c *gin.Context) {
	var r userInfo
	if err := c.Bind(&r); err != nil {
		err1 := errcode.New(errcode.ValidationError, err)
		err1.Add(err)
		api.SendResponse(c, err1, nil)
		return
	}

	jwt, err := service.GetUserJWT(r.Username, r.Password)
	if err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, &loginResp{jwt})
}

// GetUser get user by path params "username"
func GetUser(c *gin.Context) {
	username := c.Param("username")

	user, err := service.GetUser(username)
	if err != nil {
		err1 := errcode.New(errcode.CustomInternalServerError, err)
		err1.Add("get user failed.")
		api.SendResponse(c, err1, nil)
		return
	}

	if user == nil {
		api.SendResponse(c, errcode.RecordNotFoundError, nil)
		return
	}

	api.SendResponse(c, nil, user)
}

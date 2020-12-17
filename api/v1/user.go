package api

import (
	"github.com/gin-gonic/gin"

	"ginner/api"
	"ginner/errcode"
	"ginner/service/v1"
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

// Login supports request ways by curl:
/*
curl -XGET "localhost:8000/login?username=windvalley&password=12345678" -v

// Content-Type: multipart/form-data; boundary=------------------------fba2e35622281fd3
curl -XGET localhost:8000/login --form 'username=windvalley' --form 'password=12345678' -v

curl -XPOST "localhost:8000/login?username=windvalley&password=12345678" -v

// Content-Type: application/x-www-form-urlencoded
curl -XPOST localhost:8000/login -d "username=windvalley&password=12345678" -v

// Content-Type: multipart/form-data; boundary=------------------------fba2e35622281fd3
curl -XPOST localhost:8000/login --form "username=windvalley" --form "password=12345678" -v

curl -XPOST localhost:8000/login -d '{"username":"windvalley","password":"12345678"}' \
-H"Content-Type:application/json" -v
*/
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

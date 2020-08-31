package user

import (
	"github.com/gin-gonic/gin"

	"use-gin/auth"
	"use-gin/config"
	"use-gin/errcode"
	"use-gin/handler"
	"use-gin/model/rdb"
)

type response struct {
	JWT string
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

curl -XPOST localhost:8000/login -d '{"username":"windvalley","password":"12345678"}' -H"Content-Type:application/json" -v
*/
func Login(c *gin.Context) {
	var u user
	if err := c.Bind(&u); err != nil {
		err1 := errcode.New(errcode.ValidationError, err)
		err1.Add(err)
		handler.SendResponse(c, err1, nil)
		return
	}

	d, err := rdb.GetUser(u.Username)
	if err != nil {
		err1 := errcode.New(errcode.UserNotFoundError, err)
		handler.SendResponse(c, err1, nil)
		return
	}

	if err := d.CheckPassword(u.Password); err != nil {
		err1 := errcode.New(errcode.PasswordIncorrectError, err)
		handler.SendResponse(c, err1, nil)
		return
	}

	secret := config.Conf().Auth.JWTSecret
	jwtLifetime := config.Conf().Auth.JWTLifetime
	jwt, err := auth.GenerateJWT(u.Username, secret, jwtLifetime)
	if err != nil {
		err1 := errcode.New(errcode.InternalServerError, err)
		handler.SendResponse(c, err1, nil)
		return
	}

	handler.SendResponse(c, nil, &response{jwt})
}

package user

import (
	"github.com/gin-gonic/gin"

	"use-gin/errcode"
	"use-gin/handler"
	"use-gin/model/rdb"
)

type createResponse struct {
	Username string `json:"username"`
}

func Create(c *gin.Context) {
	var r user
	if err := c.Bind(&r); err != nil {
		err1 := errcode.New(errcode.ValidationError, err)
		err1.Add(err)
		handler.SendResponse(c, err1, nil)
		return
	}

	u := &rdb.User{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.EncryptPassword(); err != nil {
		err1 := errcode.New(errcode.InternalServerError, err)
		err1.Add(err)
		handler.SendResponse(c, err1, nil)
		return
	}

	if err := u.Create(); err != nil {
		err1 := errcode.New(errcode.InternalServerError, err)
		err1.Add(err)
		handler.SendResponse(c, err1, nil)
		return
	}

	handler.SendResponse(c, nil, &createResponse{r.Username})
}

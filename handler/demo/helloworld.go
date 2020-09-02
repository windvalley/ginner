package demo

import (
	"github.com/gin-gonic/gin"

	"use-gin/handler"
)

// HelloWorld a handler demo
func HelloWorld(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	handler.SendResponse(c, nil, user+", Hello world!")
	//err := errcode.New(errcode.ArgsValueError, errors.New("a system error"))
	//handler.SendResponse(c, err, nil)
}

package demo1

import (
	"github.com/gin-gonic/gin"

	"use-gin/handler"
)

func HelloWorld(c *gin.Context) {
	handler.SendResponse(c, nil, "Hello world!")
	//err := errcode.New(errcode.ArgsValueError, errors.New("a system error"))
	//handler.SendResponse(c, err, nil)
}

package demo1

import (
	"github.com/gin-gonic/gin"

	"use-gin/handler"
)

func HelloWorld(c *gin.Context) {
	handler.SendResponse(c, nil, "Hello world!")
}

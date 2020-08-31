package signdemo

import (
	"github.com/gin-gonic/gin"

	"use-gin/handler"
)

// Hello a handler demo
func Hello(c *gin.Context) {
	handler.SendResponse(c, nil, "hello world!")
}

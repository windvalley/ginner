package apiv1

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"use-gin/api"
)

// BasicAuthDemo basic auth demo
func BasicAuthDemo(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	a := 1
	b := 0
	d := a / b

	fmt.Println(d)
	api.SendResponse(c, nil, user+", hello!")
}

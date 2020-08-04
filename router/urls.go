package router

import (
	"github.com/gin-gonic/gin"
)

func urls(router *gin.Engine) {
	// The first arg is the path of the request,
	// and the second arg is the real file path in the server.
	// i.e.: request path: /s/js/xxx.js vs real path: html/statics/js/xxx.js
	router.Static("s", "html/statics")

	// url group1
	demo1 := router.Group("/v1/demo1")
	{
		demo1.GET("")
	}
}

package router

import (
	"github.com/gin-gonic/gin"

	"use-gin/handler/demo1"
	"use-gin/handler/user"
	"use-gin/midware"
)

func urls(router *gin.Engine) {
	// The first arg is the path of the request,
	// and the second arg is the real file path in the server.
	// i.e.: request path: /s/js/xxx.js vs real path: html/statics/js/xxx.js
	router.Static("s", "html/statics")

	// get jwt
	router.POST("/login", user.Login)
	router.POST("/auth", user.Login)
	router.GET("/login", user.Login)
	router.GET("/auth", user.Login)

	// group1
	g1 := router.Group("/v1/users")
	g1.POST("", user.Create) // user register request do not use jwt
	g1.Use(midware.JWT())    // use jwt
	{
		g1.GET("/:username", user.GetUser)
	}

	// group2
	g2 := router.Group("/v1/demo1")
	g2.Use(midware.JWT())
	{
		g2.GET("/eg-handlekafka", demo1.HandleKafkaDemo)
		g2.POST("/eg-handleinfluxdb", demo1.HandleInfluxdbDemo)
		g2.GET("/hello", demo1.HelloWorld)
	}
}

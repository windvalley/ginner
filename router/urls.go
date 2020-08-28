package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"use-gin/handler/demo1"
	"use-gin/handler/signdemo"
	"use-gin/handler/user"
	"use-gin/midware"
)

func urls(router *gin.Engine) {
	// The first arg is the path of the request,
	// and the second arg is the real file path in the server.
	// i.e.: request path: /s/js/xxx.js vs real path: html/statics/js/xxx.js
	router.Static("s", "html/statics")

	// for monitor the server
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// get jwt token
	router.POST("/login", user.Login)
	router.GET("/login", user.Login)

	// user manage demo
	g1 := router.Group("/v1/users")
	g1.POST("", user.Create)          // user register request can not use jwt
	g1.Use(midware.JWT())             // use jwt
	g1.Use(midware.TrafficLimit(100)) // limit 100 requests/second per source IP
	{
		g1.GET("/:username", user.GetUser)
	}

	// api signature demo
	g2 := router.Group("/v1/signdemo")
	//g2.Use(midware.Md5Sign())
	//g2.Use(midware.AESSign())
	//g2.Use(midware.RSASign())
	//g2.Use(midware.HmacMd5Sign())
	//g2.Use(midware.HmacSha1Sign())
	//g2.Use(midware.HmacSha256Sign())

	// NOTE: need to issue appKey and appSecret to users in advance
	g2.Use(midware.JWT())
	{
		g2.GET("/hello", signdemo.Hello)
		g2.POST("/hello", signdemo.Hello)
	}

	// another demo
	g3 := router.Group("/v1/demo1")
	g3.Use(midware.JWT())
	{
		g3.GET("/eg-handlekafka", demo1.HandleKafkaDemo)
		g3.POST("/eg-handleinfluxdb", demo1.HandleInfluxdbDemo)
		g3.GET("/hello", demo1.HelloWorld)
	}
}

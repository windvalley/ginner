package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"use-gin/api/apiv1"
	"use-gin/midware"
)

func urls(router *gin.Engine) {
	// The first arg is the path of the request,
	// and the second arg is the real file path in the server.
	// i.e.: request path: /s/js/xxx.js vs real path: html/statics/js/xxx.js
	router.Static("s", "html/statics")

	// For monitor the server
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Login, and get jwt token
	router.POST("/login", apiv1.Login)
	router.GET("/login", apiv1.Login)

	// User manage demo
	g1 := router.Group("/v1/users")
	g1.POST("", apiv1.CreateUser) // user register request can not use jwt
	g1.Use(midware.JWT())         // use jwt
	g1.Use(midware.UserAudit())   // enable user audit
	{
		g1.GET("/:username", apiv1.GetUser)
		g1.POST("/:username", apiv1.GetUser)
	}

	// API signature demo
	g2 := router.Group("/v1/sign-demo")
	//g2.Use(midware.APISign(midware.SignTypeMd5))
	//g2.Use(midware.APISign(midware.SignTypeAES))
	//g2.Use(midware.APISign(midware.SignTypeRSA))
	//g2.Use(midware.APISign(midware.SignTypeHmacMd5))
	//g2.Use(midware.APISign(midware.SignTypeHmacSHA1))
	//g2.Use(midware.APISign(midware.SignTypeHmacSHA256))
	// NOTE: need to issue appKey and appSecret to users in advance
	g2.Use(midware.JWT())
	{
		g2.GET("", apiv1.SignatureDemo)
	}

	// Basic auth demo
	g3 := router.Group("/v1/basic-auth-demo")
	// If necessary, we could get username in handler function by follows code line:
	// user := c.MustGet(gin.AuthUserKey).(string)
	g3.Use(gin.BasicAuth(gin.Accounts{
		"foo":   "bar",
		"admin": "123456",
	}))
	{
		g3.GET("", apiv1.BasicAuthDemo)
	}

	// handle dbs demo
	g4 := router.Group("/v1/handle-dbs-demo")
	{
		g4.GET("/kafka", apiv1.HandleKafkaDemo)
		g4.POST("/influxdb", apiv1.HandleInfluxdbDemo)
		g4.GET("/mongodb", apiv1.HandleMongodbDemo)
	}
}

package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginner/api/v1"
	apiv2 "ginner/api/v2"
	"ginner/midware"
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
	router.POST("/login", api.Login)
	router.GET("/login", api.Login)

	// User manage demo
	g1 := router.Group("/v1/users")
	g1.POST("", api.CreateUser) // user register request can not use jwt
	g1.Use(midware.JWT())       // use jwt
	{
		g1.GET("/:username", api.GetUser)
		g1.POST("/:username", api.GetUser)
	}

	// User manage version control demo
	g1V2 := router.Group("/v2/users")
	g1V2.POST("", apiv2.CreateUser) // user register request can not use jwt
	g1V2.Use(midware.JWT())         // use jwt
	{
		g1V2.GET("/:username", apiv2.GetUser)
		g1V2.POST("/:username", apiv2.GetUser)
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
		g2.GET("", api.SignatureDemo)
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
		g3.GET("", api.BasicAuthDemo)
	}

	// handle dbs demo
	g4 := router.Group("/v1/handle-dbs-demo")
	{
		g4.GET("/kafka", api.HandleKafkaDemo)
		g4.POST("/influxdb", api.HandleInfluxdbDemo)
		g4.GET("/mongodb", api.HandleMongodbDemo)
		g4.GET("/elasticsearch", api.FilterRecordsFromES)
	}
}

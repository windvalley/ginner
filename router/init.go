package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"use-gin/config"
	"use-gin/logger"
	"use-gin/midware"
)

func RouterGroup() {
	runmode := config.Config().Runmode

	switch runmode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		panic("unknown runmode")
	}

	router := gin.New()
	if runmode == "debug" {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())
	router.Use(logger.AccessLogger())
	router.Use(midware.ACL())
	router.Use(midware.RequestId())

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 Page: The API route is not correct.")
	})

	urls(router)

	if err := router.Run(config.Config().ServerPort); err != nil {
		logger.Log.Errorf("router started failed: %+v", err)
	}
}

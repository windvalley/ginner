package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"use-gin/config"
	"use-gin/logger"
	"use-gin/midware"
)

// Group router group
func Group() {
	runmode := config.Conf().Runmode

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
	router.Use(midware.RequestID())
	router.Use(midware.AccessLogger())
	router.Use(midware.ACL())
	router.Use(midware.CORS())
	router.Use(midware.GlobalTrafficLimiter(100000))
	// requests/second per client IP
	router.Use(midware.UserTrafficLimiter(100))

	router.NoRoute(func(c *gin.Context) {
		c.String(
			http.StatusNotFound,
			"404 page: url path is not correct",
		)
	})

	if runmode == "debug" {
		// NOTE: swag init first
		router.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	urls(router)

	// normal start server
	//if err := router.Run(config.Conf().ServerPort); err != nil {
	//logger.Log.Errorf("router started failed: %+v", err)
	//}

	// graceful restart or shutdown server
	server := endless.NewServer(config.Conf().ServerPort, router)
	server.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		logger.Log.Infof("current pid is %d", pid)
		if err := createPidFile(pid); err != nil {
			logger.Log.Fatalf("create pid file failed: %v", err)
		}
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Log.Warnf("%+v", err)
	}
}

func createPidFile(pid int) error {
	abPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	logDir := abPath + "/logs/"
	filename := filepath.Base(os.Args[0])
	pidFile := logDir + filename + ".pid"

	file, err := os.OpenFile(pidFile, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(strconv.Itoa(pid)); err != nil {
		return err
	}
	return nil
}

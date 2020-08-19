package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	"use-gin/config"
	"use-gin/logger"
	"use-gin/midware"
)

func RouterGroup() {
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
	router.Use(midware.AccessLogger())
	router.Use(midware.ACL())
	router.Use(midware.RequestId())

	router.NoRoute(func(c *gin.Context) {
		c.String(
			http.StatusNotFound,
			"404 Page: The API route is not correct",
		)
	})

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

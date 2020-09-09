package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"use-gin/config"
	_ "use-gin/docs" // for swagger
	"use-gin/logger"
	"use-gin/midware"
	"use-gin/util"
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
		// pprof usage:
		// Follow command will be duration for 30 seconds by default,
		// and we can benchmark the url that we want to pprof during this time.
		// go tool pprof localhost:8000/debug/pprof/profile
		// (pprof) help
		// (pprof) top 20
		// (pprof) svg
		// Subcommand svg generated report in profile001.svg in current dir,
		// and we can open profile001.svg on browser by double clicking it.
		pprof.Register(router)

		// NOTE: execute `swag init` in project root dir after updating docs.
		// path: /doc/index.html
		router.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// Live reloading the server in development stage for coding in high efficiency.
		// NOTE: Do not run server in this way: `go run main.go`,
		// and the correct way: `go build -o appname && RUNENV=dev ./appname`,
		// or `go build && export RUNENV=dev && ./service.sh start`
		go util.LiveReloadServer("./", true, []string{
			"logs",
		})
	}

	urls(router)

	// graceful restart or shutdown server
	serverPort := config.Conf().ServerPort
	server := endless.NewServer(serverPort, router)
	server.BeforeBegin = func(add string) {
		beforeServerStart(serverPort)
	}

	if config.Conf().HTTPS.Enable {
		beforeServerStart(serverPort)
		if err := server.ListenAndServeTLS(
			config.Conf().HTTPS.Cert, config.Conf().HTTPS.Key); err != nil {
			logger.Log.Warn(err)
		}
	} else {
		if err := server.ListenAndServe(); err != nil {
			logger.Log.Warn(err)
		}
	}
}

func beforeServerStart(serverPort string) {
	pid := syscall.Getpid()
	logger.Log.Debugf("current pid is %d", pid)
	logger.Log.Debugf("server port is %s", serverPort)
	if err := createPidFile(pid); err != nil {
		logger.Log.Fatalf("create pid file failed: %v", err)
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

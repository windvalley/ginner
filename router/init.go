package router

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"ginner/config"
	_ "ginner/docs" // for swagger
	"ginner/logger"
	"ginner/midware"
	"ginner/util"
)

// Group routes
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

	router.Use(midware.Recover())
	router.Use(midware.RequestID())
	router.Use(midware.UserAudit())
	router.Use(midware.AccessLogger())
	router.Use(midware.ACL())
	router.Use(midware.CORS())
	router.Use(midware.GlobalTrafficLimiter(100000))
	router.Use(midware.UserTrafficLimiter(100)) // requests/second per client IP

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
		// and the correct way: `go build -o appname && ./appname -c conf/dev.config.conf`
		go util.LiveReloadServer([]string{
			"logs",
		})
	}

	urls(router)

	// Check whether the service has been started successfully
	go pingServer()

	// graceful restart or shutdown server
	serverPort := config.Conf().ServerPort
	server := endless.NewServer(serverPort, router)
	server.BeforeBegin = func(add string) {
		beforeServerStart(serverPort, runmode)
	}

	var err error
	if config.Conf().HTTPS.Enable {
		beforeServerStart(serverPort, runmode)
		err = server.ListenAndServeTLS(config.Conf().HTTPS.Cert, config.Conf().HTTPS.Key)
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		if runmode == "debug" {
			fmt.Printf("%s %v\n", color.FgYellow.Render("[Endless-warning]"), err)
		} else {
			logger.Log.Warn(err)
		}
	}
}

func beforeServerStart(serverPort, runmode string) {
	pid := syscall.Getpid()

	if runmode == "debug" {
		fmt.Printf("%s current pid is %s\n",
			color.FgCyan.Render("[Endless-debug]"), color.FgGreen.Render(pid))
		fmt.Printf("%s server port is %s\n",
			color.FgCyan.Render("[Endless-debug]"), color.FgGreen.Render(serverPort))
	}

	if err := createPidFile(pid); err != nil {
		logger.Log.Fatalf("[Endless-fatal] create pid file failed: %v", err)
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

func pingServer() {
	pid := syscall.Getpid()

	schema := "http://"
	if config.Conf().HTTPS.Enable {
		schema = "https://"
	}

	pingURL := schema + "127.0.0.1" + config.Conf().ServerPort + "/ping"

	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	logger.Log.Debugf("checking url: %s", pingURL)

	for i := 0; i < 6; i++ {
		resp, err := http.Get(pingURL)
		if err == nil && resp.StatusCode == 200 {
			logger.Log.Debugf("server(%d) started", pid)
			return
		}

		time.Sleep(time.Second)
	}

	logger.Log.Fatalf("server(%d) start failed", pid)
}

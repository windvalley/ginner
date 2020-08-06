package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"use-gin/config"
)

var Log = logrus.New()

func Init() {
	abPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	logDir := abPath + "/logs"
	if !IsPathExist(logDir) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			panic(err)
		}
	}

	Log.SetLevel(logrus.DebugLevel)

	accessLog := path.Join(logDir, "access.log")
	accesslogWiter, err := rotatelogs.New(
		accessLog+"-%Y-%m-%d",
		rotatelogs.WithLinkName(accessLog),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		panic(fmt.Sprintf("log rotate faield: %v", err))
	}

	errorLog := path.Join(logDir, "error.log")
	errorlogWiter, err := rotatelogs.New(
		errorLog+"-%Y-%m-%d",
		rotatelogs.WithLinkName(errorLog),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		panic(fmt.Sprintf("log rotate faield: %v", err))
	}

	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: accesslogWiter,
		logrus.InfoLevel:  errorlogWiter,
		logrus.WarnLevel:  errorlogWiter,
		logrus.ErrorLevel: errorlogWiter,
		logrus.PanicLevel: errorlogWiter,
		logrus.FatalLevel: errorlogWiter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Log.AddHook(lfHook)

	// Only log in file, not to screen(io.Stderr) in production.
	if config.Config().Runmode == "release" {
		Log.Out = ioutil.Discard
	}
}

func AccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latencyTime := end.Sub(start).Seconds()
		reqPath := c.Request.URL.Path
		clientIP := c.ClientIP()
		reqMethod := c.Request.Method
		statusCode := c.Writer.Status()

		Log.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqPath,
		}).Debug()
	}
}

// IsPathExist file or dir is exist or not
func IsPathExist(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// InitCMDLog preserve error log for cmd
func InitCMDLog() {
	abPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	logDir := abPath + "/logs"
	if !IsPathExist(logDir) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			panic(err)
		}
	}

	Log.SetLevel(logrus.DebugLevel)

	filename := filepath.Base(os.Args[0])
	logfile := logDir + "/" + filename + ".log"
	logWiter, err := rotatelogs.New(
		logfile+"-%Y-%m-%d",
		rotatelogs.WithLinkName(logfile),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		panic(fmt.Sprintf("log rotate faield: %v", err))
	}

	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: logWiter,
		logrus.InfoLevel:  logWiter,
		logrus.WarnLevel:  logWiter,
		logrus.ErrorLevel: logWiter,
		logrus.PanicLevel: logWiter,
		logrus.FatalLevel: logWiter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Log.AddHook(lfHook)
	//Log.Out = ioutil.Discard
}

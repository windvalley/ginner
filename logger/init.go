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
	logDir, err := createLogdir()
	if err != nil {
		panic(fmt.Sprintf("create log dir failed: %v", err))
	}

	Log.SetLevel(logrus.DebugLevel)

	accessLog := path.Join(logDir, "access.log")
	accesslogWriter, err := getLogWriter(accessLog)
	if err != nil {
		panic(fmt.Sprintf("log rotate failed: %v", err))
	}

	errorLog := path.Join(logDir, "error.log")
	errorlogWriter, err := getLogWriter(errorLog)
	if err != nil {
		panic(fmt.Sprintf("log rotate failed: %v", err))
	}

	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: accesslogWriter,
		logrus.InfoLevel:  errorlogWriter,
		logrus.WarnLevel:  errorlogWriter,
		logrus.ErrorLevel: errorlogWriter,
		logrus.PanicLevel: errorlogWriter,
		logrus.FatalLevel: errorlogWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Log.AddHook(lfHook)

	// Only log in file, not to screen(io.Stderr) in production.
	if config.Conf().Runmode == "release" {
		Log.Out = ioutil.Discard
	}
}

func AccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latencyTime := time.Since(start).Seconds()
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

// InitCMDLog preserve error log for cmd
func InitCMDLog() {
	logDir, err := createLogdir()
	if err != nil {
		panic(fmt.Sprintf("create log dir failed: %v", err))
	}

	Log.SetLevel(logrus.DebugLevel)

	filename := filepath.Base(os.Args[0])
	logfile := logDir + "/" + filename + ".log"
	logWriter, err := getLogWriter(logfile)
	if err != nil {
		panic(fmt.Sprintf("log rotate failed: %v", err))
	}

	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.FatalLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Log.AddHook(lfHook)
	//Log.Out = ioutil.Discard
}

func createLogdir() (string, error) {
	abPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	logDir := abPath + "/" + config.Conf().Log.Dirname
	if !IsPathExist(logDir) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			return "", err
		}
	}

	return logDir, nil
}

func getLogWriter(logFile string) (*rotatelogs.RotateLogs, error) {
	rotationHours := time.Duration(config.Conf().Log.RotationHours)
	saveDays := time.Duration(config.Conf().Log.SaveDays)

	logWiter, err := rotatelogs.New(
		logFile+"-%Y-%m-%d",
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithRotationTime(rotationHours*time.Hour),
		rotatelogs.WithMaxAge(saveDays*24*time.Hour),
	)
	if err != nil {
		return nil, err
	}

	return logWiter, nil
}

// IsPathExist file or dir is exist or not
func IsPathExist(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

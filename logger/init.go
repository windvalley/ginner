package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"ginner/config"
)

// Log logger instance
var Log = logrus.New()

const timeFormat = "2006-01-02 15:04:05"

// Config as logger init parameters
type Config struct {
	Runmode       string
	Dirname       string
	LogFormat     string
	LogLevel      string
	RotationHours time.Duration
	SaveDays      time.Duration
}

// Init logger
func Init() {
	log := config.Conf().Log

	logDir, err := createLogdir(log.Dirname)
	if err != nil {
		panic(fmt.Sprintf("create log dir '%s' failed: %s", logDir, err))
	}

	Log.SetLevel(getLogLevel(log.LogLevel))

	accessLog := path.Join(logDir, "access.log")
	accesslogWriter, err := getLogWriter(accessLog, log.RotationHours, log.SaveDays)
	if err != nil {
		panic(err)
	}

	errorLog := path.Join(logDir, "error.log")
	errorlogWriter, err := getLogWriter(errorLog, log.RotationHours, log.SaveDays)
	if err != nil {
		panic(err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  accesslogWriter,
		logrus.DebugLevel: errorlogWriter,
		logrus.WarnLevel:  errorlogWriter,
		logrus.ErrorLevel: errorlogWriter,
		logrus.FatalLevel: errorlogWriter,
		logrus.PanicLevel: errorlogWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: timeFormat,
	})

	if log.LogFormat == "txt" {
		lfHook = lfshook.NewHook(writeMap, &logrus.TextFormatter{
			TimestampFormat: timeFormat,
		})
	}

	Log.AddHook(lfHook)

	// Only write log to file, not to screen(io.Stderr) in production.
	if config.Conf().Runmode == "release" {
		Log.Out = ioutil.Discard
	}
}

// InitCmd init logger for subproject(in cmd/) that using config file of main project
func InitCmd(logConf *Config) {
	logDir, err := createLogdir(logConf.Dirname)
	if err != nil {
		panic(fmt.Sprintf("create log dir '%s' failed: %s", logDir, err))
	}

	Log.SetLevel(getLogLevel(logConf.LogLevel))

	filename := filepath.Base(os.Args[0])
	logfile := logDir + "/" + filename + ".log"
	logWriter, err := getLogWriter(logfile, logConf.RotationHours, logConf.SaveDays)
	if err != nil {
		panic(err)
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
		TimestampFormat: timeFormat,
	})

	if logConf.LogFormat == "txt" {
		lfHook = lfshook.NewHook(writeMap, &logrus.TextFormatter{
			TimestampFormat: timeFormat,
		})
	}

	Log.AddHook(lfHook)

	if logConf.Runmode == "release" {
		Log.Out = ioutil.Discard
	}
}

func createLogdir(dirName string) (string, error) {
	abPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	logDir := abPath + "/" + dirName
	if !isPathExist(logDir) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			return logDir, err
		}
	}

	return logDir, nil
}

func getLogWriter(logFile string, rotationHours, saveDays time.Duration) (*rotatelogs.RotateLogs, error) {
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

func getLogLevel(level string) (logLevel logrus.Level) {
	switch level {
	case "trace":
		logLevel = logrus.TraceLevel
	case "debug":
		logLevel = logrus.DebugLevel
	case "warn":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	case "fatal":
		logLevel = logrus.FatalLevel
	case "panic":
		logLevel = logrus.PanicLevel
	default:
		panic("no such log level: " + level)
	}

	return logLevel
}

func isPathExist(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

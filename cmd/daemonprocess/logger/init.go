package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"ginner/cmd/daemonprocess/cfg"
)

const (
	log1FileName = "log1.log"
	log2FileName = "log2.log"
)

var (
	// Log1 log1.log
	Log1 = logrus.New()
	// Log2 log2.log
	Log2 = logrus.New()
)

const timeFormat = "2006-01-02 15:04:05"

// Init logger
func Init() {
	logCfg := cfg.Conf().Log

	logDir, err := createLogdir(logCfg.Dirname)
	if err != nil {
		panic(fmt.Sprintf("create log dir failed: %v", err))
	}

	Log1.SetLevel(logrus.DebugLevel)
	Log2.SetLevel(logrus.DebugLevel)

	log1File := logDir + "/" + log1FileName
	log2File := logDir + "/" + log2FileName

	log1LfsHook, err := getLfsHook(log1File, logCfg.LogFormat, logCfg.RotationHours, logCfg.SaveDays)
	if err != nil {
		panic(err)
	}

	log2LfsHook, err := getLfsHook(log2File, logCfg.LogFormat, logCfg.RotationHours, logCfg.SaveDays)
	if err != nil {
		panic(err)
	}

	Log1.AddHook(log1LfsHook)
	Log2.AddHook(log2LfsHook)

	if cfg.Conf().Runmode == "release" {
		Log1.Out = ioutil.Discard
		Log2.Out = ioutil.Discard
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
			return "", err
		}
	}

	return logDir, nil
}

func getLfsHook(logFile, logFormat string,
	rotationHours, saveDays int) (*lfshook.LfsHook, error) {
	logWriter, err := rotatelogs.New(
		logFile+"-%Y-%m-%d",
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithRotationTime(time.Duration(rotationHours)*time.Hour),
		rotatelogs.WithMaxAge(time.Duration(saveDays)*24*time.Hour),
	)
	if err != nil {
		return nil, err
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

	if logFormat == "txt" {
		lfHook = lfshook.NewHook(writeMap, &logrus.TextFormatter{
			TimestampFormat: timeFormat,
		})
	}

	return lfHook, nil
}

func isPathExist(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

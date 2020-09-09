package util

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
	"use-gin/logger"
)

var (
	colorGreen = string([]byte{27, 91, 57, 55, 59, 51, 50, 59, 49, 109})
	colorRed   = string([]byte{27, 91, 57, 55, 59, 51, 49, 59, 49, 109})
	colorReset = string([]byte{27, 91, 48, 109})
)

// LiveReloadServer Auto build and graceful restart the server while file changed.
// Mainly for development stage.
func LiveReloadServer(rootPath string, monitorAllFiles bool, excludeDirs []string) {
	startTime := time.Now()

	for {
		filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
			if path == ".git" && info.IsDir() {
				return filepath.SkipDir
			}

			for _, x := range excludeDirs {
				if x == path {
					return filepath.SkipDir
				}
			}

			// ignore hidden files
			if filepath.Base(path)[0] == '.' {
				return nil
			}

			if (monitorAllFiles || filepath.Ext(path) == ".go") && info.ModTime().After(startTime) {
				scanCallback(path)
				startTime = time.Now()
				return errors.New("done")
			}

			return nil
		})
		time.Sleep(500 * time.Millisecond)
	}
}

func scanCallback(path string) {
	pid := os.Getpid()
	logger.Log.Warnf(
		"%s[LiveReloadServer]%s %s%s%s has been changed, and start to reload the server(%d)...",
		colorGreen,
		colorReset,
		colorRed,
		path,
		colorReset,
		pid,
	)

	gobuild := exec.Command("/usr/local/bin/go", "build")
	if err := gobuild.Start(); err != nil {
		logger.Log.Errorf(
			"%s[LiveReloadServer]%s go build start failed: %v",
			colorGreen,
			colorReset,
			err,
		)
		return
	}
	if err := gobuild.Wait(); err != nil {
		logger.Log.Errorf(
			"%s[LiveReloadServer]%s go build wait failed: %v",
			colorGreen,
			colorReset,
			err,
		)

		return
	}

	reload := exec.Command("kill", "-SIGHUP", strconv.Itoa(pid))
	if err := reload.Start(); err != nil {
		logger.Log.Errorf(
			"%s[LiveReloadServer]%s kill -SIGHUP %d start failed: %v",
			colorGreen,
			colorReset,
			pid,
			err,
		)
		return
	}
	if err := reload.Wait(); err != nil {
		logger.Log.Errorf(
			"%s[LiveReloadServer]%s kill -SIGHUP %d wait failed: %v",
			colorGreen,
			colorReset,
			pid,
			err,
		)
		return
	}
}

package util

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gookit/color"
)

// LiveReloadServer Auto build and graceful restart the server while file changed,
// and this feature is mainly for development stage.
func LiveReloadServer(excludeDirs []string) {
	startTime := time.Now()
	rootPath := "."

	for {
		filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
			if isSkipPath(path, info, excludeDirs) {
				return filepath.SkipDir
			}

			// ignore hidden files
			if filepath.Base(path)[0] == '.' {
				return nil
			}

			if info.ModTime().After(startTime) {
				buildAndReload(path)
				startTime = time.Now()
				return errors.New("done")
			}

			return nil
		})

		time.Sleep(500 * time.Millisecond)
	}
}

func isSkipPath(path string, info os.FileInfo, excludeDirs []string) bool {
	if path == ".git" && info.IsDir() {
		return true
	}

	for _, x := range excludeDirs {
		if x == path {
			return true
		}
	}

	return false
}

func buildAndReload(path string) {
	pid := os.Getpid()

	red := color.FgRed.Render
	green := color.FgGreen.Render

	fmt.Printf(
		"%s %s has been changed, and begin to reload the server(%d)...\n",
		green("[LiveReloadServer-debug]"),
		red(path),
		pid,
	)

	if err := gobuild(); err != nil {
		fmt.Printf("%s %v\n", red("[LiveReloadServer-error]"), err)
		return
	}

	if err := reloadServer(pid); err != nil {
		fmt.Printf("%s %v\n", red("[LiveReloadServer-error]"), err)
	}
}

func gobuild() error {
	cmd := exec.Command("go", "build")

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("go build command start error: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("go build command wait error: %v", err)
	}

	return nil
}

func reloadServer(pid int) error {
	if err := syscall.Kill(pid, syscall.SIGHUP); err != nil {
		return fmt.Errorf("reload server(%d) error: %v", pid, err)
	}

	return nil
}

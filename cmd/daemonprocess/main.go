package main

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"ginner/config"
	"ginner/util"

	"ginner/cmd/daemonprocess/cfg"
	"ginner/cmd/daemonprocess/logger"
)

func init() {
	config.InitCmd(&cfg.Config)

	logger.Init()
}

func main() {
	abPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	pidDir := abPath + "/" + cfg.Conf().Log.Dirname
	lock, lockFile, err := util.ProcessLock(pidDir)
	if err != nil {
		logger.Log1.Fatal(err)
	}
	defer os.Remove(lockFile)
	defer lock.Close()

	ctx, cancel := context.WithCancel(context.Background())
	go yourLogic(ctx)

	sigC := make(chan os.Signal)
	signal.Notify(sigC)
	for sig := range sigC {
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
			syscall.SIGKILL:
			logger.Log2.Infof("%v signal captured, quit.", sig)
			cancel()
			os.Remove(lockFile)
			os.Exit(1)
		}
	}
}

func yourLogic(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:

			// your specific logic

		}
	}
}

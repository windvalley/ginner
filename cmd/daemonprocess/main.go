package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"ginner/config"
	"ginner/logger"
	"ginner/util"

	"ginner/cmd/daemonprocess/cfg"
)

func init() {
	config.InitCmd(&cfg.Config)

	logger.InitCmdLogger(
		cfg.Conf().Log.Dirname,
		cfg.Conf().Log.LogFormat,
		cfg.Conf().Log.RotationHours,
		cfg.Conf().Log.SaveDays,
	)
}

func main() {
	lock, lockFile, err := util.ProcessLock()
	if err != nil {
		logger.Log.Fatal(err)
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
			logger.Log.Infof("%v signal captured, quit.", sig)
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

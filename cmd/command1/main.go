package main

import (
	"fmt"
	"os"

	"use-gin/logger"
	"use-gin/utils"
)

func main() {
	lock, lockFile, err := utils.ProcessLock()
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer os.Remove(lockFile)
	defer lock.Close()

	fmt.Println("vim-go")
}

package main

import (
	"os"
	"path/filepath"

	"ginner/config"
	"ginner/db/es"
	"ginner/db/rdb"
	"ginner/db/redclus"
	"ginner/logger"
	"ginner/util"
)

func init() {
	// use config file of main project
	config.Init()

	logger.Init()

	redclus.Init()

	es.Init()
}

func main() {
	abPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	pidDir := abPath + "/" + config.Conf().Log.Dirname
	lock, lockFile, err := util.ProcessLock(pidDir)
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer os.Remove(lockFile)
	defer lock.Close()

	rdb.Init()
	defer rdb.Close()

	data, err := getFinalData()
	if err != nil {
		logger.Log.Fatalf("get final domain records data failed: %s", err)
	}

	engine := newDataSaveEngine("indexname")
	if err := engine.GetOldIndices().SaveData(data).Clean().Err; err != nil {
		logger.Log.Errorf("sync data to es failed: %v", err)
	}
}

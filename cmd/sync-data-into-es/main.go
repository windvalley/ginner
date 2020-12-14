package main

import (
	"os"

	"ginner/config"
	"ginner/db/es"
	"ginner/db/rdb"
	"ginner/db/redclus"
	"ginner/logger"
	"ginner/util"
)

func init() {
	config.LoadFromCLIParams()

	logger.InitCMDLog()

	redclus.Init()

	es.Init()
}

func main() {
	lock, lockFile, err := util.ProcessLock()
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

package main

import (
	//"ginner/model/influxdb"
	//"ginner/model/kafka"

	"ginner/config"
	"ginner/cron"
	"ginner/db/mongodb"
	"ginner/db/rdb"
	"ginner/db/redis"
	"ginner/logger"
	"ginner/router"
	"ginner/util"
)

func init() {
	// Load config from command line parameters.
	//    e.g. ./ginner -c conf/dev.config.conf
	//config.LoadFromCLIParams()

	// Load config from system environment variable RUNENV: prod/dev
	//    e.g. RUNENV=dev ./ginner
	//config.LoadFromENV()

	// If load config from CLI params failed,
	// then load config from system environment variable RUNENV,
	// and the value of RUNENV can only be dev or prod.
	config.Init()

	logger.Init()

	cron.Init()

	// cache
	redis.Init()
	//redclus.Init()
}

// @title Use-Gin API
// @version 0.1.0
// @description Using Go Gin to develop high quality applications(Web API) efficiently.
// @contact.name Windvalley
// @contact.email i@sre.im
// @license.name MIT
// @license.url https://github.com/windvalley/ginner/blob/master/LICENSE
// @host ginner.sre.im:8000
// @BasePath /api
func main() {
	// relation db
	rdb.Init()
	defer rdb.Close()

	// migrate relation db tables
	util.MigrateRDBTables()

	// mongodb
	mongodb.Init()
	defer mongodb.Close()

	// influxdb
	//influxdb.Init()
	//defer influxdb.Close()

	// kafka
	//kafka.InitConsumer()
	//kafka.InitProducer()

	router.Group()
}

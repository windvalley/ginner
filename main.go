package main

import (
	//"ginner/model/influxdb"
	//"ginner/model/kafka"

	"ginner/config"
	"ginner/cron"
	"ginner/db/mongodb"
	"ginner/db/rdb"
	"ginner/logger"
	"ginner/router"
	"ginner/util"
)

func init() {
	// Load config file from command line parameters.
	//    e.g. ./ginner -c conf/dev.config.conf
	config.Init()

	logger.Init()

	cron.Init()

	// cache
	//redis.Init()
	// or
	//redclus.Init()

	// search
	//es.Init()
}

// @title ginner API
// @version 0.0.1
// @description For developing high quality applications(Web API) efficiently by Go Gin.
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

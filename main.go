package main

import (
	"use-gin/config"
	"use-gin/cron"
	"use-gin/logger"
	"use-gin/model/rdb"
	"use-gin/model/redclus"
	"use-gin/model/redis"
	"use-gin/router"
)

func init() {
	// Load config from command line parameters.
	//    e.g. ./use-gin -c conf/dev.config.conf
	//config.LoadFromCLIParams()

	// Load config from system environment variable RUNENV: prod/dev
	//    e.g. RUNENV=dev ./use-gin
	//config.LoadFromENV()

	// If load config from CLI params failed,
	// then load config from system environment variable RUNENV,
	// and the value of RUNENV can only be dev or prod.
	config.Init()

	logger.Init()

	cron.Init()

	redis.Init()

	redclus.Init()
}

// @title Use-Gin API
// @version 0.1.0
// @description Using Go Gin to develop high quality applications(Web API) efficiently.
// @contact.name Windvalley
// @contact.email i@sre.im
// @license.name MIT
// @license.url https://github.com/windvalley/use-gin/blob/master/LICENSE
// @host use-gin.sre.im:8000
// @BasePath /api
func main() {
	// relation db
	rdb.Init()
	rdb.DBs.MySQL.AutoMigrate(
		&rdb.User{},
	)
	defer rdb.Close()

	// influxdb
	//influxdb.Init()
	//defer influxdb.Close()

	// kafka
	//kafka.InitKafkaConsumer()
	//kafka.InitKafkaProducer()

	router.Group()
}

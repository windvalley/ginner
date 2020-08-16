package main

import (
	"use-gin/config"
	"use-gin/cron"
	"use-gin/logger"
	"use-gin/router"
)

func init() {
	// load config from command line parameters.
	//    e.g. ./use-gin -c conf/dev.config.conf
	//config.LoadFromCLIParams()

	// load from system environment variable RUNENV: prod/dev
	//    e.g. export RUNENV=dev
	// ./use-gin
	config.LoadFromENV()

	// logger
	logger.Init()

	// cron
	cron.Init()

	// relation db
	//rdb.Init()
	//rdb.DBs.MySQL.Set("gorm:table_options", "ENGIN=InnoDB").AutoMigrate(
	////&rdb.xxx{},
	//)
	//defer rdb.Close()

	// influxdb
	//influxdb.Init()
	//defer influxdb.Close()

	// kafka
	//kafka.InitKafkaConsumer()
	//kafka.InitKafkaProducer()
}

func main() {
	router.RouterGroup()
}

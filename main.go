package main

import (
	"use-gin/config"
	"use-gin/cron"
	"use-gin/logger"
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
}

func main() {
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

	router.RouterGroup()
}

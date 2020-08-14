package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"

	"use-gin/config"
	"use-gin/cron"
	"use-gin/logger"
	"use-gin/model/kafka"
	"use-gin/model/mysql"
	"use-gin/router"
)

func main() {
	// config
	cfg := pflag.StringP("config", "c", "", "Specify your configuration file")
	pflag.Parse()
	if *cfg == "" {
		binName := filepath.Base(os.Args[0])
		fmt.Printf("missing parameter\nUsage of %s:\n  -c, --config string"+
			"   Specify your configuration file\n", binName)
		os.Exit(2)
	}
	config.ParseConfig(*cfg)

	// logger
	logger.Init()

	// mysqldb
	mysql.DBs.Init()
	mysql.DBs.Useraccount.Set("gorm:table_options", "ENGIN=InnoDB").AutoMigrate(
	//&mysql.xxx{},
	)
	defer mysql.DBs.Close()

	// kafka
	kafka.InitKafkaClient()
	kafka.InitKafkaProducer()

	// cron
	cron.Init()

	// router
	router.RouterGroup()
}

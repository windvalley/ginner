package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"

	"use-gin/config"
	"use-gin/logger"
	"use-gin/model/mysql"
	"use-gin/router"
)

func main() {
	cfg := pflag.StringP("config", "c", "", "Specify your configuration file")
	pflag.Parse()
	if *cfg == "" {
		binName := filepath.Base(os.Args[0])
		fmt.Printf("missing parameter\nUsage of %s:\n  -c, --config string"+
			"   Specify your configuration file\n", binName)
		os.Exit(2)
	}

	config.ParseConfig(*cfg)

	logger.Init()

	mysql.DBs.Init()
	mysql.DBs.Useraccount.Set("gorm:table_options", "ENGIN=InnoDB").AutoMigrate(
	//&mysql.xxx{},
	)
	defer mysql.DBs.Close()

	router.RouterGroup()
}

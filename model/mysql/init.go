package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"use-gin/config"
	"use-gin/logger"
)

type Databases struct {
	Useraccount *gorm.DB
}

var DBs *Databases

func (*Databases) Init() {
	DBs = &Databases{
		Useraccount: GetDemoDB(),
	}
}

func (*Databases) Close() {
	DBs.Useraccount.Close()
}

func GetDemoDB() *gorm.DB {
	address := config.Config().DBDemo.Address
	dbname := config.Config().DBDemo.DBName
	user := config.Config().DBDemo.User
	password := config.Config().DBDemo.Password

	return Connect(user, password, address, dbname)
}

func Connect(username, password, address, dbname string) *gorm.DB {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			username,
			password,
			address,
			dbname))
	if err != nil {
		//logger.Log.Fatalf("Connect mysql database %s failed: %v", dbname, err)
		logger.Log.Panicf("Connect mysql database %s failed: %v", dbname, err)
	}

	if config.Config().Runmode != "release" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)

	return db
}

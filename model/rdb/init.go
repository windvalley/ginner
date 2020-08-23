package rdb

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"

	"use-gin/config"
	"use-gin/logger"
)

// if have any other databases, just put it in this struct
type Databases struct {
	MySQL      *gorm.DB
	PostgreSQL *gorm.DB
}

var DBs *Databases

func Init() {
	DBs = &Databases{
		MySQL: GetMySQL(),
		//PostgreSQL: GetPostgreSQL(),
	}
}

func Close() {
	DBs.MySQL.Close()
	//DBs.PostgreSQL.Close()
}

func GetMySQL() *gorm.DB {
	dbtype := config.Conf().MySQL.DBType
	address := config.Conf().MySQL.Address
	dbname := config.Conf().MySQL.DBName
	user := config.Conf().MySQL.User
	password := config.Conf().MySQL.Password

	return Connect(dbtype, user, password, address, dbname)
}

func GetPostgreSQL() *gorm.DB {
	dbtype := config.Conf().PostgreSQL.DBType
	address := config.Conf().PostgreSQL.Address
	dbname := config.Conf().PostgreSQL.DBName
	user := config.Conf().PostgreSQL.User
	password := config.Conf().PostgreSQL.Password

	return Connect(dbtype, user, password, address, dbname)
}

func Connect(dbtype, username, password, address, dbname string) *gorm.DB {
	db, err := gorm.Open(dbtype,
		fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			username,
			password,
			address,
			dbname,
		))
	if err != nil {
		logger.Log.Fatalf("connect mysql database %s failed: %v", dbname, err)
	}

	if config.Conf().Runmode != "release" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db
}

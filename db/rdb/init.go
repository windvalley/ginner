package rdb

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"

	//_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mssql"

	"use-gin/config"
	"use-gin/logger"
)

// Databases if have any other databases, just put it in this struct
type Databases struct {
	MySQL      *gorm.DB
	PostgreSQL *gorm.DB
}

// DBs client instance of dbs
var DBs *Databases

// Init DBs initialization
func Init() {
	DBs = &Databases{
		MySQL: GetDBConn("mysql"),
		//PostgreSQL: GetDBConn("postgresql),
	}
}

// Close close dbs connection
func Close() {
	DBs.MySQL.Close()
	//DBs.PostgreSQL.Close()
}

// GetDBConn get a db instance of relation databases
func GetDBConn(db string) *gorm.DB {
	dbtype := config.Conf().RDBs[db].DBType
	address := config.Conf().RDBs[db].Address
	dbname := config.Conf().RDBs[db].DBName
	user := config.Conf().RDBs[db].User
	password := config.Conf().RDBs[db].Password
	maxIdleConns := config.Conf().RDBs[db].MaxIdleConns
	maxOpenConns := config.Conf().RDBs[db].MaxOpenConns

	return Connect(dbtype, user, password, address, dbname,
		maxIdleConns, maxOpenConns)
}

// Connect relation db connect
func Connect(dbtype, username, password, address, dbname string,
	maxIdleConns, maxOpenConns int) *gorm.DB {
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
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxIdleConns)

	return db
}

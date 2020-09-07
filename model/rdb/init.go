package rdb

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"

	//_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"

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
		MySQL: GetMySQL(),
		//PostgreSQL: GetPostgreSQL(),
	}

	autoMigrateTables()
}

// Close close dbs connection
func Close() {
	DBs.MySQL.Close()
	//DBs.PostgreSQL.Close()
}

// GetMySQL get db instance of mysql
func GetMySQL() *gorm.DB {
	dbtype := config.Conf().RDBs["mysql"].DBType
	address := config.Conf().RDBs["mysql"].Address
	dbname := config.Conf().RDBs["mysql"].DBName
	user := config.Conf().RDBs["mysql"].User
	password := config.Conf().RDBs["mysql"].Password
	maxIdleConns := config.Conf().RDBs["mysql"].MaxIdleConns
	maxOpenConns := config.Conf().RDBs["mysql"].MaxOpenConns

	return Connect(dbtype, user, password, address, dbname,
		maxIdleConns, maxOpenConns)
}

// GetPostgreSQL get db instance of postgresql
func GetPostgreSQL() *gorm.DB {
	dbtype := config.Conf().RDBs["postgresql"].DBType
	address := config.Conf().RDBs["postgresql"].Address
	dbname := config.Conf().RDBs["postgresql"].DBName
	user := config.Conf().RDBs["postgresql"].User
	password := config.Conf().RDBs["postgresql"].Password
	maxIdleConns := config.Conf().RDBs["postgresql"].MaxIdleConns
	maxOpenConns := config.Conf().RDBs["postgresql"].MaxOpenConns

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

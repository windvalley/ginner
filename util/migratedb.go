package util

import (
	"use-gin/db/rdb"
	"use-gin/model"
)

// MigrateRDBTables migrate relation db's schemas and keep schemas up to date
func MigrateRDBTables() {
	rdb.DBs.MySQL.AutoMigrate(
		&model.User{},
		&model.UserOperationLog{},
	)
}

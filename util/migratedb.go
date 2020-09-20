package util

import (
	"use-gin/db/rdb"
	"use-gin/model"
)

// MigrateDBTables migrate relation db's schemas and keep schemas up to date
func MigrateDBTables() {
	rdb.DBs.MySQL.AutoMigrate(
		&model.User{},
		&model.UserOperationLog{},
	)
}

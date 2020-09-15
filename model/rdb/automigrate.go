package rdb

func autoMigrateTables() {
	DBs.MySQL.AutoMigrate(
		&User{},
		&UserOperationLog{},
	)
}

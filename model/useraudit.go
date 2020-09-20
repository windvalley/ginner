package model

import (
	"time"
	"use-gin/db/rdb"
)

// UserOperationLog user operation audit
type UserOperationLog struct {
	Model
	Username   string
	ClientIP   string
	ReqMethod  string
	ReqPath    string
	ReqBody    string
	ReqReferer string
	UserAgent  string
	ReqTime    time.Time
	ReqLatency float64
	HTTPStatus int
	ResCode    string
	ResMessage string
	ResData    string
}

// Create insert a record
func (u *UserOperationLog) Create() error {
	return rdb.DBs.MySQL.Create(&u).Error
}

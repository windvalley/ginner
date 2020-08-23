package rdb

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"create_at" time_format:"2006-01-02"`
	UpdatedAt time.Time  `json:"update_at" time_format:"2006-01-02"`
	DeletedAt *time.Time `sql:"index" json:"delete_at" time_format:"2006-01-02"`
}

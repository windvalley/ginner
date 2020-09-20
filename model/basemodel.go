package model

import "time"

// Model base model
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"create_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time  `json:"update_at" time_format:"2006-01-02 15:04:05"`
	DeletedAt *time.Time `sql:"index" json:"delete_at" time_format:"2006-01-02 15:04:05"`
}

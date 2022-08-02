package model

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagID        int64
	TagName      string `gorm:"varchar(128);unique"`
	Introduction string `gorm:"varchar(256);unique"`
}

type TagInfo struct {
	TagID   int64  `json:"tag_id"`
	TagName string `json:"tag_name"`
}

type TagDetail struct {
	TagID        int64     `json:"tag_id"`
	TagName      string    `json:"tag_name"`
	Introduction string    `json:"introduction"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
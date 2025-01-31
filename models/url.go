package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	ShortCode   string `gorm:"unique;not null"`
	Url         string `gorm:"not null"`
	AccessCount uint   `gorm:"default:0"`
}

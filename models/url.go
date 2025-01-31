package models

import (
	"time"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	OriginalURL string `gorm:"not null" json:"url"`
	ShortCode   string `gorm:"unique;not null" json:"shortCode"`
	AccessCount int    `gorm:"default:0" json:"accessCount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
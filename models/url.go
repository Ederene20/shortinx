package models

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	UniqueID string `json:"unique_id" gorm:"unique"`
	LongUrl string `json:"long_url"`
	ShortUrl string `json:"short_url" gorm:"unique"`
}
package models

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID               uint `gorm:"primaryKey"`
	Link             string
	ShortendLink     string `gorm:",unique"`
	ShortendFullLink string `gorm:",unique"`
	TimesVisited     uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

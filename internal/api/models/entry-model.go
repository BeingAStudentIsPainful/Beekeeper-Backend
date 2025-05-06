package models

import (
	"beekeeper-backend/internal/types"
	"time"
)

type Entry struct {
	ID        uint `gorm:"primaryKey"`
	HiveID    uint
	HiveName  string          `gorm:"not null"`
	Content   string          `gorm:"not null"`
	Type      types.EntryType `gorm:"not null"`
	CreatedAt time.Time
}

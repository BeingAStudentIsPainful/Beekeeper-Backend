package models

import "time"

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	HiveID    uint   `gorm:"not null"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
}

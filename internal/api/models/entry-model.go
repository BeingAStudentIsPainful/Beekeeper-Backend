package models

import "time"

type EntryType string

const (
	EntryTypeLog  EntryType = "log"
	EntryTypeTask EntryType = "task"
)

type Entry struct {
	ID        uint   `gorm:"primaryKey"`
	HiveID    uint   `gorm:"not null"`
	Content   string `gorm:"not null"`
	Type      EntryType
	CreatedAt time.Time
}

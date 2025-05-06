package models

import "time"

type Hive struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	QueenYear   int
	Location    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Entries     []Entry `gorm:"constraint:OnDelete:CASCADE"`
}

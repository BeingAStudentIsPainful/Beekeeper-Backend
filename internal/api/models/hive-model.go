package models

import "time"

type Hive struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Logs      []Log  `gorm:"constraint:OnDelete:CASCADE"`
	Tasks     []Task `gorm:"constraint:OnDelete:CASCADE"`
}

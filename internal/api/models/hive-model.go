package models

import "time"

// Hive represents a beehive in the management system
type Hive struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	HiveName  int       `json:"hive_name" gorm:"not null" example:"123"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-15T10:30:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-15T10:30:00Z"`
	Logs      []Log     `json:"logs,omitempty" gorm:"foreignKey:HiveID;constraint:OnDelete:CASCADE"`
	Tasks     []Task    `json:"tasks,omitempty" gorm:"foreignKey:HiveID;constraint:OnDelete:CASCADE"`
}

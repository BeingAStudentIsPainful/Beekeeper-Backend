package models

import "time"

type Task struct {
    ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
    HiveID    int       `json:"hive_id" gorm:"not null" example:"123"`
    Content   string    `json:"content" gorm:"not null" example:"Check honey levels and replace frames"`
    CreatedAt time.Time `json:"created_at" example:"2024-01-15T10:30:00Z"`
    UpdatedAt time.Time `json:"updated_at" example:"2024-01-15T10:30:00Z"`
}

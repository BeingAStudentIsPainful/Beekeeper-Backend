package config

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/database"
	"beekeeper-backend/internal/utils"
)

func Initialization() {
	utils.LoadEnv()
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Hive{}, &models.Log{})
}

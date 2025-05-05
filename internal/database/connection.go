package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("../internal/database/beekeeper-database.db"), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		return
	}

	DB = database
	log.Println("Database connection established")
}
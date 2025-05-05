package main

import (
	"beekeeper-backend/internal/database"
	"beekeeper-backend/internal/utils"
	"fmt"
	"os"
)

func main() {
	utils.LoadEnv()
	database.ConnectDatabase()

	port := os.Getenv("PORT")
	fmt.Println(port)
}
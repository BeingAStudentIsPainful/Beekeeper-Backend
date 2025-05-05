package main

import (
	"beekeeper-backend/internal/config"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Initialization()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := gin.Default()

	app.Run(":" + port)
}

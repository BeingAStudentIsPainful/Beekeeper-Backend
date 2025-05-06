package main

import (
	"beekeeper-backend/internal/api/routes"
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
	api := app.Group("/api")

	routes.EntryRoutes(api)

	app.Run(":" + port)
}

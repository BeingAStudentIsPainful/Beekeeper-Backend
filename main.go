package main

import (
	"beekeeper-backend/docs"
	"beekeeper-backend/internal/api/routes"
	"beekeeper-backend/internal/config"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Beekeeper API
// @version         1.0
// @description     A beekeeping management API built with Go and Gin framework
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

func main() {
	db := config.Initialization()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := gin.Default()

	// Swagger documentation route
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := app.Group("/api")
	routes.TaskRoutes(api, db)
	routes.LogRoutes(api, db)
	routes.HiveRoutes(api, db)

	// Programmatically set swagger info
	docs.SwaggerInfo.Title = "Beekeeper API"
	docs.SwaggerInfo.Description = "A beekeeping management API built with Go and Gin framework"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app.Run(":" + port)
}

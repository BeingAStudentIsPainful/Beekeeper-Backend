package routes

import (
	"beekeeper-backend/internal/api/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TaskRoutes(app *gin.RouterGroup, db *gorm.DB) {
	handler := handlers.TaskHandler{DB: db}

	taskRoutes := app.Group("/tasks")

	taskRoutes.POST("/", handler.CreateTask)
	taskRoutes.GET("/:id", handler.GetTaskByID)

}

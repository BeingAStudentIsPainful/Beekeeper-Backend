package routes

import (
	"beekeeper-backend/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func EntryRoutes(app *gin.RouterGroup) {
	entryRoutes := app.Group("/tasks")

	entryRoutes.POST("/", handlers.CreateTask)
	entryRoutes.GET("/:id", handlers.GetTaskByID)

}

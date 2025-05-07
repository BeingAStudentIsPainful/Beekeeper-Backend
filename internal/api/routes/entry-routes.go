package routes

import (
	"beekeeper-backend/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func EntryRoutes(app *gin.RouterGroup) {
	entryRoutes := app.Group("/entries")

	entryRoutes.POST("/", handlers.CreateEntry)
	entryRoutes.GET("/:id", handlers.GetEntryByID)

}

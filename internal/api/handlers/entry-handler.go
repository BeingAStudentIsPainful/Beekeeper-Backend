package handlers

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/database"
	"beekeeper-backend/internal/types"

	"github.com/gin-gonic/gin"
)

func CreateEntry(c *gin.Context) {
	var input types.CreateEntryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Fill all required fields",
		})
		return
	}

	entry := models.Entry{
		Content:  input.Content,
		Type:     input.Type,
		HiveName: input.HiveName,
	}

	if err := database.DB.Create(&entry).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create entry",
		})
		return
	}

	c.JSON(201, entry)
}

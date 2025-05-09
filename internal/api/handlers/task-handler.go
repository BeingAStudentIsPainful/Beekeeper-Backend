package handlers

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/database"
	"beekeeper-backend/internal/types"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var input types.CreateEntryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Fill all required fields",
		})
		return
	}

	task := models.Task{
		HiveID:  input.HiveID,
		Content: input.Content,
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create entry",
		})
		return
	}

	c.JSON(201, task)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Task not found",
		})
	}

	c.JSON(200, gin.H{
		"data": task,
	})
}

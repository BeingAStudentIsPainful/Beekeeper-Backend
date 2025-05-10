package handlers

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/types"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	DB *gorm.DB
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input types.CreateEntryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	task := models.Task{
		HiveID:  input.HiveID,
		Content: input.Content,
	}

	if err := h.DB.Create(&task).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(201, gin.H{"data": task})
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := h.DB.First(&task, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(200, gin.H{"data": task})
}

func GetAllTasks(c *gin.Context) {

}

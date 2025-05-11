package handlers

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	BaseHandler
}

// Create task
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

// Get task by id
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	var task models.Task
	h.GetByID(c, &task)
}

// Get all tasks
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	var tasks []models.Task

	if err := h.DB.Find(&tasks).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	c.JSON(200, gin.H{"data": tasks})
}

// Update task
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := h.DB.First(&task, id).Error; err != nil {
		c.JSON(500, gin.H{"errror": "Task not found"})
		return
	}

	var input types.UpdateEntryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if input.Content != nil {
		task.Content = *input.Content
	}
	if input.HiveID != nil {
		task.HiveID = *input.HiveID
	}

	if err := h.DB.Save(&task).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(200, gin.H{"data": task})
}

// Delete task
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var task models.Task
	if err := h.DB.First(&task, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	if err := h.DB.Delete(&task, id).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete task"})
		return
	}

	c.Status(204)
}

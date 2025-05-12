package handlers

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/types"

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

	h.CreateEntry(c, &task)
}

// Get task by id
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	var task models.Task
	h.GetEntryByID(c, &task)
}

// Get all tasks
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	var tasks []models.Task
	h.GetAllEntries(c, &tasks)
}

// Update task
func (h *TaskHandler) UpdateTask(c *gin.Context) {

	var task models.Task
	var input types.UpdateEntryInput
	h.UpdateEntry(c, &task, &input, func(model any, input any) {
		t := model.(*models.Task)
		i := input.(*types.UpdateEntryInput)

		if i.Content != nil {
			t.Content = *i.Content
		}
		if i.HiveID != nil {
			t.HiveID = *i.HiveID
		}
	})
}

// Delete task
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	var task models.Task
	h.DeleteEntry(c, &task)
}

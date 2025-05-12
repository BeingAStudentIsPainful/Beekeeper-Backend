package handlers

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/types"

	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	BaseHandler
}

// Create log
func (h *LogHandler) CreateLog(c *gin.Context) {
	var input types.CreateEntryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	log := models.Log{
		HiveID:  input.HiveID,
		Content: input.Content,
	}

	h.CreateEntry(c, &log)
}

// Get log by id
func (h *LogHandler) GetLogByID(c *gin.Context) {
	var log models.Log
	h.GetEntryByID(c, &log)
}

// Get all logs
func (h *LogHandler) GetAllLogs(c *gin.Context) {
	var logs []models.Log
	h.GetAllEntries(c, &logs)
}

// Update log
func (h *LogHandler) UpdateLog(c *gin.Context) {

	var log models.Log
	var input types.UpdateEntryInput
	h.UpdateEntry(c, &log, &input, func(model any, input any) {
		t := model.(*models.Log)
		i := input.(*types.UpdateEntryInput)

		if i.Content != nil {
			t.Content = *i.Content
		}
		if i.HiveID != nil {
			t.HiveID = *i.HiveID
		}
	})
}

// Delete log
func (h *LogHandler) DeleteLog(c *gin.Context) {
	var log models.Log
	h.DeleteEntry(c, &log)
}

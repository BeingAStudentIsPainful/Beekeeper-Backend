package handlers

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/types"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LogHandler handles all CRUD operations for logs.
type LogHandler struct {
	BaseHandler
}

// CreateLog handles POST /logs.
// @Summary Create a new log entry
// @Description Create a new log entry for a hive. If the hive doesn't exist, it will be created automatically.
// @Tags logs
// @Accept json
// @Produce json
// @Param log body types.CreateEntryInput true "Log creation data"
// @Success 201 {object} map[string]models.Log "Successfully created log"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /logs [post]
func (h *LogHandler) CreateLog(c *gin.Context) {
	var input types.CreateEntryInput
	var hive models.Hive

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	log := models.Log{
		HiveID:  input.HiveID,
		Content: input.Content,
	}

	// Check if the hive exists by HiveName
	if err := h.DB.Where("hive_name = ?", log.HiveID).First(&hive).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create hive remotely if not found
			newHive, err := h.CreateHiveRemote(c, input.HiveID)
			if err != nil {
				c.JSON(500, gin.H{"error": "Could not create hive"})
				return
			}
			hive = *newHive
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	h.CreateEntry(c, &log)
}

// GetLogByID handles GET /logs/:id.
// @Summary Get a log entry by ID
// @Description Retrieve a specific log entry by its ID
// @Tags logs
// @Accept json
// @Produce json
// @Param id path int true "Log ID"
// @Success 200 {object} map[string]models.Log "Successfully retrieved log"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 404 {object} map[string]string "Log not found"
// @Router /logs/{id} [get]
func (h *LogHandler) GetLogByID(c *gin.Context) {
	var log models.Log
	h.GetEntryByID(c, &log)
}

// GetAllLogs handles GET /logs.
// @Summary Get all log entries
// @Description Retrieve all log entries from the database
// @Tags logs
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]models.Log "Successfully retrieved all logs"
// @Failure 500 {object} map[string]string "Failed to retrieve logs"
// @Router /logs [get]
func (h *LogHandler) GetAllLogs(c *gin.Context) {
	var logs []models.Log
	h.GetAllEntries(c, &logs)
}

// GetLastLog handles GET /logs/last
// @Summary Get the most recent log entry
// @Description Retrieve the last log entry based on creation time
// @Tags logs
// @Accept json
// @Produce json
// @Success 200 {object} map[string]models.Log "Successfully retrieved last log"
// @Failure 404 {object} map[string]string "No logs found"
// @Router /logs/last [get]
func (h *LogHandler) GetLastLog(c *gin.Context) {
	var log models.Log
	h.GetLastEntry(c, &log)
}

// UpdateLog handles PUT /logs/:id.
// @Summary Update a log entry
// @Description Update an existing log entry by ID
// @Tags logs
// @Accept json
// @Produce json
// @Param id path int true "Log ID"
// @Param log body types.UpdateEntryInput true "Log update data"
// @Success 200 {object} map[string]models.Log "Successfully updated log"
// @Failure 400 {object} map[string]string "Invalid ID or input"
// @Failure 404 {object} map[string]string "Log not found"
// @Failure 500 {object} map[string]string "Failed to update log"
// @Router /logs/{id} [put]
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

// DeleteLog handles DELETE /logs/:id.
// @Summary Delete a log entry
// @Description Delete a log entry by ID
// @Tags logs
// @Accept json
// @Produce json
// @Param id path int true "Log ID"
// @Success 204 "Successfully deleted log"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 404 {object} map[string]string "Log not found"
// @Router /logs/{id} [delete]
func (h *LogHandler) DeleteLog(c *gin.Context) {
	var log models.Log
	h.DeleteEntry(c, &log)
}

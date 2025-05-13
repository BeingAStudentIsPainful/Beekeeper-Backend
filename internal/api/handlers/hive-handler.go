package handlers

import (
	"beekeeper-backend/internal/api/models"

	"github.com/gin-gonic/gin"
)

type HiveHandler struct {
	BaseHandler
}

func (h *HiveHandler) CreateHive(c *gin.Context) {
	var hive models.Hive

	if err := c.ShouldBindBodyWithJSON(&hive); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&hive).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": hive})
}
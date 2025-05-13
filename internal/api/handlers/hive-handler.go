package handlers

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/types"

	"github.com/gin-gonic/gin"
)

type HiveHandler struct {
	BaseHandler
}

func (h *HiveHandler) CreateHive(c *gin.Context) {
	var input types.CreateHiveInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	hive := models.Hive{
		HiveName: input.HiveName,
	}

	if err := h.DB.Create(&hive).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not create hive"})
		return
	}

	c.JSON(201, gin.H{"data": hive})
}

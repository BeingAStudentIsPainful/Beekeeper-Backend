package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseHandler struct {
	DB *gorm.DB
}

func (h *BaseHandler) GetByID(c *gin.Context, model any) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.DB.First(&model, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Entry not found"})
		return
	}

	c.JSON(200, gin.H{"data": model})
}

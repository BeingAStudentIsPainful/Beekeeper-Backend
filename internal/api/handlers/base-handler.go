package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseHandler struct {
	DB *gorm.DB
}

func (h *BaseHandler) CreateEntry(c *gin.Context, model any) {
	if err := h.DB.Create(model).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create entry"})
		return
	}

	c.JSON(201, gin.H{"data": model})
}

func (h *BaseHandler) GetAllEntries(c *gin.Context, model any) {
	if err := h.DB.Find(model).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve entry"})
		return
	}

	c.JSON(200, gin.H{"data": model})
}

func (h *BaseHandler) GetEntryByID(c *gin.Context, model any) {
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

func (h *BaseHandler) UpdateEntry(
	c *gin.Context,
	model any,
	input any,
	applyChanges func(any, any)) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.DB.First(&model, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Entry not found"})
		return
	}

	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	applyChanges(model, input)

	if err := h.DB.Save(model).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to save"})
		return
	}

	c.JSON(200, gin.H{"data": model})
}

func (h *BaseHandler) DeleteEntry(c *gin.Context, model any) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.DB.First(&model, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Entry not found"})
		return
	}

	if err := h.DB.Delete(&model, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Entry not found"})
		return
	}

	c.Status(204)
}

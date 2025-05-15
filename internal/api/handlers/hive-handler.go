package handlers

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/types"

	"github.com/gin-gonic/gin"
)

// HiveHandler provides CRUD operations for hive-related data.
type HiveHandler struct {
	BaseHandler
}

// CreateHive handles POST /hives.
// It creates a new hive based on JSON input.
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

// GetAllHives handles GET /hives.
// It retrieves and returns all hives from the database.
func (h *HiveHandler) GetAllHives(c *gin.Context){
	var hives []models.Hive
	h.GetAllEntries(c, &hives)
}

// GetHiveByID handles GET /hives/:id.
// It retrieves a single hive by its ID.
func (h *HiveHandler) GetHiveByID(c *gin.Context){
	var hive models.Hive
	h.GetEntryByID(c, &hive)
}

// UpdateHive handles PUT /hives/:id.
// It updates a hive using provided JSON input.
func (h *HiveHandler) UpdateHive(c *gin.Context){
	var hive models.Hive
	var input types.UpdateHiveInput
	h.UpdateEntry(c, &hive, &input, func(model any, input any){
		h := model.(*models.Hive)
		i := input.(*types.UpdateHiveInput)

		if i.HiveName != nil{
			h.HiveName = *i.HiveName
		}
	})
}


// DeleteHive handles DELETE /hives/:id.
// It deletes a hive by its ID.
func (h *HiveHandler) DeleteHive(c *gin.Context){
	var hive models.Hive
	h.DeleteEntry(c, &hive)
}
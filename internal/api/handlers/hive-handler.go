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

func (h *HiveHandler) GetAllHives(c *gin.Context){
	var hives []models.Hive
	h.GetAllEntries(c, &hives)
}

func (h *HiveHandler) GetHiveByID(c *gin.Context){
	var hive models.Hive
	h.GetEntryByID(c, &hive)
}

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

func (h *HiveHandler) DeleteHive(c *gin.Context){
	var hive models.Hive
	h.DeleteEntry(c, &hive)
}
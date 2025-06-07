package handlers

import (
	"beekeeper-backend/internal/api/models"
	"beekeeper-backend/internal/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HiveHandler provides CRUD operations for hive-related data.
type HiveHandler struct {
    BaseHandler
}

// CreateHive handles POST /hives.
// @Summary      Create a new hive
// @Description  Create a new hive with the provided information
// @Tags         hives
// @Accept       json
// @Produce      json
// @Param        hive  body      types.CreateHiveInput  true  "Hive data"
// @Success      201   {object}  map[string]interface{}  "Successfully created hive"
// @Failure      400   {object}  map[string]string       "Invalid input"
// @Failure      500   {object}  map[string]string       "Could not create hive"
// @Router       /hives [post]
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
// @Summary      List all hives
// @Description  Get a list of all hives
// @Tags         hives
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "List of hives"
// @Failure      500  {object}  map[string]string       "Failed to retrieve hives"
// @Router       /hives [get]
func (h *HiveHandler) GetAllHives(c *gin.Context) {
    var hives []models.Hive
    h.GetAllEntries(c, &hives)
}

// GetHiveByID handles GET /hives/:id.
// @Summary      Get hive by ID
// @Description  Get a single hive by its hive name/ID
// @Tags         hives
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Hive ID"
// @Success      200  {object}  map[string]interface{}  "Hive details"
// @Failure      400  {object}  map[string]string       "Invalid ID"
// @Failure      404  {object}  map[string]string       "Hive not found"
// @Router       /hives/{id} [get]
func (h *HiveHandler) GetHiveByID(c *gin.Context) {
    var hive models.Hive
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.DB.Where("hive_name = ?", id).First(&hive).Error; err != nil {
        c.JSON(404, gin.H{"error": "Entry not found"})
        return
    }

    c.JSON(200, gin.H{"data": hive})
}

// UpdateHive handles PATCH /hives/:id.
// @Summary      Update hive
// @Description  Update an existing hive by its ID
// @Tags         hives
// @Accept       json
// @Produce      json
// @Param        id    path      int                     true  "Hive ID"
// @Param        hive  body      types.UpdateHiveInput   true  "Updated hive data"
// @Success      200   {object}  map[string]interface{}  "Updated hive"
// @Failure      400   {object}  map[string]string       "Invalid ID or input"
// @Failure      404   {object}  map[string]string       "Hive not found"
// @Failure      500   {object}  map[string]string       "Failed to save"
// @Router       /hives/{id} [patch]
func (h *HiveHandler) UpdateHive(c *gin.Context) {
    var hive models.Hive
    var input types.UpdateHiveInput
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.DB.Where("hive_name = ?", id).First(&hive).Error; err != nil {
        c.JSON(404, gin.H{"error": "Entry not found"})
        return
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    hive.HiveName = *input.HiveName

    if err := h.DB.Save(&hive).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to save"})
        return
    }

    c.JSON(200, gin.H{"data": hive})
}

// DeleteHive handles DELETE /hives/:id.
// @Summary      Delete hive
// @Description  Delete a hive by its ID
// @Tags         hives
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "Hive ID"
// @Success      204  "Successfully deleted"
// @Failure      400  {object}  map[string]string  "Invalid ID"
// @Failure      404  {object}  map[string]string  "Hive not found"
// @Router       /hives/{id} [delete]
func (h *HiveHandler) DeleteHive(c *gin.Context) {
    var hive models.Hive
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.DB.Where("hive_name = ?", id).First(&hive).Error; err != nil {
        c.JSON(404, gin.H{"error": "Entry not found"})
        return
    }

    if err := h.DB.Delete(&hive).Error; err != nil {
        c.JSON(404, gin.H{"error": "Failed to delete hive"})
        return
    }

    c.Status(204)
}
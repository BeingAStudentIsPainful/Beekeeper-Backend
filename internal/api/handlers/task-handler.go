package handlers

import (
    "beekeeper-backend/internal/api/models"
    "beekeeper-backend/internal/types"
    "errors"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// TaskHandler handles all CRUD operations related to tasks.
type TaskHandler struct {
    BaseHandler
}

// CreateTask handles POST /tasks.
// @Summary Create a new task
// @Description Create a new task for a hive. If the hive doesn't exist, it will be created automatically.
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body types.CreateEntryInput true "Task creation data"
// @Success 201 {object} map[string]models.Task "Successfully created task"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
    var input types.CreateEntryInput
    var hive models.Hive

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    task := models.Task{
        HiveID:  input.HiveID,
        Content: input.Content,
    }

    // Check if hive exists
    if err := h.DB.Where("hive_name = ?", task.HiveID).First(&hive).Error; err != nil {
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

    h.CreateEntry(c, &task)
}

// GetTaskByID handles GET /tasks/:id.
// @Summary Get a task by ID
// @Description Retrieve a specific task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]models.Task "Successfully retrieved task"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 404 {object} map[string]string "Task not found"
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
    var task models.Task
    h.GetEntryByID(c, &task)
}

// GetAllTasks handles GET /tasks.
// @Summary Get all tasks
// @Description Retrieve all tasks from the database
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]models.Task "Successfully retrieved all tasks"
// @Failure 500 {object} map[string]string "Failed to retrieve tasks"
// @Router /tasks [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
    var tasks []models.Task
    h.GetAllEntries(c, &tasks)
}

// GetLastTask handles GET /tasks/last
// @Summary Get the most recent task
// @Description Retrieve the last task based on creation time
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} map[string]models.Task "Successfully retrieved last task"
// @Failure 404 {object} map[string]string "No tasks found"
// @Router /tasks/last [get]
func (h *TaskHandler) GetLastTask(c *gin.Context) {
    var task models.Task
    h.GetLastEntry(c, &task)
}

// UpdateTask handles PUT /tasks/:id.
// @Summary Update a task
// @Description Update an existing task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body types.UpdateEntryInput true "Task update data"
// @Success 200 {object} map[string]models.Task "Successfully updated task"
// @Failure 400 {object} map[string]string "Invalid ID or input"
// @Failure 404 {object} map[string]string "Task not found"
// @Failure 500 {object} map[string]string "Failed to update task"
// @Router /tasks/{id} [put]
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

// DeleteTask handles DELETE /tasks/:id.
// @Summary Delete a task
// @Description Delete a task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 204 "Successfully deleted task"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 404 {object} map[string]string "Task not found"
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
    var task models.Task
    h.DeleteEntry(c, &task)
}
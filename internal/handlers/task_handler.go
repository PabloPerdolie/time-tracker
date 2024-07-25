package handlers

import (
	"EffectiveMobileTestTask/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// taskHandler handles task-related requests
type taskHandler struct {
	taskService services.TasksService
}

// NewTaskHandler creates a new taskHandler
func NewTaskHandler(t services.TasksService) TaskHandler {
	return &taskHandler{
		taskService: t,
	}
}

// GetTasksByUser
// @Summary Get tasks by user ID
// @Description Get a list of tasks for a specific user within an optional date range
// @Tags tasks
// @Produce json
// @Param id path int true "User ID"
// @Param start_date query string false "Start date in YYYY-MM-DD format"
// @Param end_date query string false "End date in YYYY-MM-DD format"
// @Success 200 {array} models.Task
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /tasks/user/{id} [get]
func (th *taskHandler) GetTasksByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	startDate := c.Query("start_date")
	if startDate == "" {
		startDate = "1970-01-01"
	}

	endDate := c.Query("end_date")
	if endDate == "" {
		endDate = time.Now().Format("2025-01-02")
	}
	tasks, err := th.taskService.GetTasksByUser(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// StartTask
// @Summary Start a new task
// @Description Start a new task with the given details
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param description body string true "Task description"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /tasks/start/{id} [post]
func (th *taskHandler) StartTask(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var description struct {
		desc string `json:"description"`
	}
	if err := c.ShouldBindJSON(&description); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := th.taskService.StartTask(userID, description.desc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// EndTask
// @Summary End a task
// @Description End an existing task with the given details
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /tasks/end/{taskId} [post]
func (th *taskHandler) EndTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := th.taskService.EndTask(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

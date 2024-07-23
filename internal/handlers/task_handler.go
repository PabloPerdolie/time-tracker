package handlers

import (
	"EffectiveMobileTestTask/internal/models"
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
		endDate = time.Now().Format("2006-01-02")
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
// @Param task body models.Task true "Task data"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /tasks/start [post]
func (th *taskHandler) StartTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.StartTime = time.Now()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	err := th.taskService.StartTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to start task"})
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
// @Param task body models.Task true "Task data"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /tasks/end [post]
func (th *taskHandler) EndTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.EndTime = time.Now()
	task.UpdatedAt = time.Now()
	task.Duration = time.Until(task.StartTime).String()

	err := th.taskService.EndTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to end task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

package handlers

import (
	"EffectiveMobileTestTask/internal/models"
	"EffectiveMobileTestTask/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type taskHandler struct {
	taskService services.TasksService
}

func NewTaskHandler(t services.TasksService) TaskHandler {
	return &taskHandler{
		taskService: t,
	}
}

func (th *taskHandler) GetTasksByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	tasks, err := th.taskService.GetTasksByUser(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

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

func (th *taskHandler) EndTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.EndTime = time.Now()
	task.UpdatedAt = time.Now()
	task.Duration = time.Until(task.StartTime)

	err := th.taskService.EndTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to end task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

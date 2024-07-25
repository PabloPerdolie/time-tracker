package services

import (
	"EffectiveMobileTestTask/internal/models"
	"EffectiveMobileTestTask/internal/repository"
)

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(t repository.TaskRepository) TasksService {
	return &taskService{
		taskRepo: t,
	}
}

func (t *taskService) GetTasksByUser(userID int, startDate, endDate string) ([]*models.Task, error) {
	return t.taskRepo.GetByUser(userID, startDate, endDate)
}

func (t *taskService) StartTask(userID int, desc string) (*models.Task, error) {
	return t.taskRepo.Create(userID, desc)
}

func (t *taskService) EndTask(taskID int) (*models.Task, error) {
	return t.taskRepo.Update(taskID)
}

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

func (t *taskService) StartTask(task *models.Task) error {
	return t.taskRepo.Create(task)
}

func (t *taskService) EndTask(task *models.Task) error {
	return t.taskRepo.Update(task)
}

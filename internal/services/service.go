package services

import "EffectiveMobileTestTask/internal/models"

type UserService interface {
	GetUsers(filter map[string]interface{}, offset, limit int) ([]*models.User, error)
	GetUserByID(id int) (*models.User, error)
	AddUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type TasksService interface {
	GetTasksByUser(userID int, startDate, endDate string) ([]*models.Task, error)
	StartTask(task *models.Task) error
	EndTask(task *models.Task) error
}

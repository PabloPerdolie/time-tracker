package repository

import (
	"EffectiveMobileTestTask/internal/models"
)

type UserRepository interface {
	GetAll(filter map[string]interface{}, offset, limit int) ([]*models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id int) error
}

type TaskRepository interface {
	GetByUser(userID int, startDate, endDate string) ([]*models.Task, error)
	Create(task *models.Task) error
	Update(task *models.Task) error
}

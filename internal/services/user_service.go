package services

import (
	"EffectiveMobileTestTask/internal/models"
	"EffectiveMobileTestTask/internal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(u repository.UserRepository) UserService {
	return &userService{
		userRepo: u,
	}
}

func (u *userService) GetUsers(filter map[string]interface{}, offset, limit int) ([]*models.User, error) {
	return u.userRepo.GetAll(filter, offset, limit)
}

func (u *userService) GetUserByID(id int) (*models.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userService) AddUser(user *models.User) error {
	return u.userRepo.Create(user)
}

func (u *userService) UpdateUser(user *models.User) error {
	return u.userRepo.Update(user)
}

func (u *userService) DeleteUser(id int) error {
	return u.userRepo.Delete(id)
}

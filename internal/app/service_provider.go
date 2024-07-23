package app

import (
	"EffectiveMobileTestTask/internal/config"
	"EffectiveMobileTestTask/internal/handlers"
	"EffectiveMobileTestTask/internal/repository"
	"EffectiveMobileTestTask/internal/repository/postgres"
	"EffectiveMobileTestTask/internal/services"
	"EffectiveMobileTestTask/internal/utils"
)

type serviceProvider struct {
	taskHandler    handlers.TaskHandler
	taskService    services.TasksService
	taskRepository repository.TaskRepository
	userHandler    handlers.UserHandler
	userService    services.UserService
	userRepository repository.UserRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) initServices() error {
	dbConfig := config.CONFIG.DB
	db, err := utils.InitDB(dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Host, dbConfig.Port)
	if err != nil {
		return err
	}
	sp.taskRepository = postgres.NewTaskRepository(db)
	sp.userRepository = postgres.NewUserRepository(db)
	sp.taskService = services.NewTaskService(sp.taskRepository)
	sp.userService = services.NewUserService(sp.userRepository)
	sp.taskHandler = handlers.NewTaskHandler(sp.taskService)
	sp.userHandler = handlers.NewUserHandler(sp.userService)
	return nil
}

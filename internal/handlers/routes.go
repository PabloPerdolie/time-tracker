package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(userHandler UserHandler, taskHandler TaskHandler) *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", userHandler.GetUsers)
		userGroup.POST("/", userHandler.AddUser)
		userGroup.PUT("/", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}

	taskGroup := r.Group("/tasks")
	{
		taskGroup.GET("/user/:id", taskHandler.GetTasksByUser)
		taskGroup.POST("/start", taskHandler.StartTask)
		taskGroup.POST("/end", taskHandler.EndTask)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

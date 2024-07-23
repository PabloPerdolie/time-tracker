package handlers

import (
	"github.com/gin-gonic/gin"
)

type TaskHandler interface {
	GetTasksByUser(c *gin.Context)
	StartTask(c *gin.Context)
	EndTask(c *gin.Context)
}

type UserHandler interface {
	GetUsers(c *gin.Context)
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

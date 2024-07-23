package handlers

import (
	"EffectiveMobileTestTask/internal/models"
	"EffectiveMobileTestTask/internal/services"
	"EffectiveMobileTestTask/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(u services.UserService) UserHandler {
	return &userHandler{
		userService: u,
	}
}

func (uh *userHandler) GetUsers(c *gin.Context) {
	filter := make(map[string]interface{})
	if surname := c.Query("surname"); surname != "" {
		filter["surname"] = surname
	}
	if name := c.Query("name"); name != "" {
		filter["name"] = name
	}

	offset := 0
	limit := 10
	users, err := uh.userService.GetUsers(filter, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uh *userHandler) AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := utils.GetPersonInfo(user.PassportSeries, user.PassportNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get person info"})
		return
	}

	user.Surname = person.Surname
	user.Name = person.Name
	user.Patronymic = person.Patronymic
	user.Address = person.Address
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = uh.userService.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *userHandler) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.UpdatedAt = time.Now()

	err := uh.userService.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *userHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uh.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

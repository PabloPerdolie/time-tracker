package handlers

import (
	"EffectiveMobileTestTask/internal/models"
	"EffectiveMobileTestTask/internal/services"
	"EffectiveMobileTestTask/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// userHandler handles user-related requests
type userHandler struct {
	userService services.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(u services.UserService) UserHandler {
	return &userHandler{
		userService: u,
	}
}

// GetUsers
// @Summary Get users
// @Description Get a list of users with optional filters
// @Tags users
// @Produce json
// @Param surname query string false "User surname"
// @Param name query string false "User name"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{} "Error"
// @Router /users [get]
func (uh *userHandler) GetUsers(c *gin.Context) {
	filter := make(map[string]interface{})
	if surname := c.Query("surname"); surname != "" {
		filter["surname"] = surname
	}
	if name := c.Query("name"); name != "" {
		filter["name"] = name
	}
	if passportNumber := c.Query("passport_number"); passportNumber != "" {
		filter["passport_number"] = passportNumber
	}
	if passportSeries := c.Query("passport_series"); passportSeries != "" {
		filter["passport_series"] = passportSeries
	}
	if patronymic := c.Query("patronymic"); patronymic != "" {
		filter["patronymic"] = patronymic
	}
	if address := c.Query("address"); address != "" {
		filter["address"] = address
	}

	var limit, offset int
	offset1 := c.Query("offset")
	if offset1 == "" {
		offset1 = "0"
	}
	offset, err := strconv.Atoi(offset1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limit1 := c.Query("limit")
	if limit1 == "" {
		limit1 = "10"
	}

	limit, err = strconv.Atoi(limit1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := uh.userService.GetUsers(filter, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

type Passport struct {
	Number string `json:"passportNumber"`
}

// AddUser
// @Summary Create a new user
// @Description Create a new user with the given details
// @Tags users
// @Accept json
// @Produce json
// @Param user body Passport true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /users [post]
func (uh *userHandler) AddUser(c *gin.Context) {
	var passport Passport
	if err := c.ShouldBindJSON(&passport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parts := strings.Split(passport.Number, " ")

	person, err := utils.GetPersonInfo(parts[0], parts[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	user.Surname = person.Surname
	user.Name = person.Name
	user.Patronymic = person.Patronymic
	user.Address = person.Address
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = uh.userService.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser
// @Summary Update an existing user
// @Description Update user details
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /users [put]
func (uh *userHandler) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.UpdatedAt = time.Now()

	err := uh.userService.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DeleteUser
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User deleted"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /users/{id} [delete]
func (uh *userHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uh.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

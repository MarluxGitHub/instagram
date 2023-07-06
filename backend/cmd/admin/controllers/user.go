package controllers

import (
	"strconv"

	"github.com/MarluxGitHub/instagram/internal/database/repository"
	"github.com/MarluxGitHub/instagram/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body User true "User info"
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "bad request"
// @Router /api/v1/user [post]
func CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	repository.CreateUser(user)

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// GetUserById godoc
// @Summary Get a user by id
// @Description Get a user by id
// @Tags user
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "bad request"
// @Router /api/v1/user/{id} [get]
func GetUserById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	// get int value of id
	intID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	user := repository.GetUserById(intID)

	c.JSON(200, gin.H{
		"message": "ok",
		"user": user,
	})
}
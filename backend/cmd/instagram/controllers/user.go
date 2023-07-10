package controllers

import (
	"github.com/MarluxGitHub/instagram/internal/database/repository"
	"github.com/MarluxGitHub/instagram/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := models.User{}

	c.BindJSON(&user)

	repository.CreateUser(user)

	c.JSON(200, gin.H{
		"message": "User created",
	})
}
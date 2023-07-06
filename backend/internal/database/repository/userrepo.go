package repository

import (
	"github.com/MarluxGitHub/instagram/internal/database"
	"github.com/MarluxGitHub/instagram/internal/models"
)

func GetUserById(id int) models.User {
	db := database.Connect()

	var user models.User

	db.First(&user, id)

	return user
}

func CreateUser(user models.User) {
	db := database.Connect()

	db.Create(&user)
}
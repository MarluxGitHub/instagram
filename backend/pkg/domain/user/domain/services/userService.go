package services

import "github.com/MarluxGitHub/instagram/pkg/domain/user/domain/models"

type UserService interface {
	GetUser(id int) *models.User
	CreateUser(*models.User) *models.User
	UpdateUser(*models.User) *models.User
	DeleteUser(id int) bool

	GetUsers() []*models.User

	GetFollowers(id int) []*models.User
	GetFollowed(id int) []*models.User
}
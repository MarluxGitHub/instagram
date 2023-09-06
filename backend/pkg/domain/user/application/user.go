package application

import (
	"github.com/MarluxGitHub/instagram/pkg/domain/user/domain/models"
	"github.com/MarluxGitHub/instagram/pkg/domain/user/domain/services"
	"github.com/MarluxGitHub/instagram/pkg/domain/user/infrastructure"
	"gorm.io/gorm"
)

type UserApplication struct {
	userService services.UserService
}

func NewUserApplication(db *gorm.DB) UserApplication {
	return UserApplication{
		userService: infrastructure.NewUserRepository(db),
	}
}

func (m *UserApplication) GetUser(id int) *models.User {
	return m.userService.GetUser(id)
}

func (m *UserApplication) CreateUser(user *models.User) *models.User {
	return m.userService.CreateUser(user)
}

func (m *UserApplication) UpdateUser(user *models.User) *models.User {
	return m.userService.UpdateUser(user)
}

func (m *UserApplication) DeleteUser(id int) bool {
	return m.userService.DeleteUser(id)
}

func (m *UserApplication) GetUsers() []*models.User {
	return m.userService.GetUsers()
}

func (m *UserApplication) GetFollowers(id int) []*models.User {
	return m.userService.GetFollowers(id)
}

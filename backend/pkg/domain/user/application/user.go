package application

import (
	"github.com/MarluxGitHub/instagram/pkg/domain/user/domain/models"
	"github.com/MarluxGitHub/instagram/pkg/domain/user/infrastructure"
	"gorm.io/gorm"
)

type UserApplication struct {
	userRepository infrastructure.UserRepository
}

func NewUserApplication(db *gorm.DB) UserApplication {
	return UserApplication{
		userRepository: infrastructure.NewUserRepository(db),
	}
}

func (m *UserApplication) GetUser(id int) *models.User {
	return m.userRepository.GetUser(id)
}

func (m *UserApplication) CreateUser(user *models.User) *models.User {
	return m.userRepository.CreateUser(user)
}

func (m *UserApplication) UpdateUser(user *models.User) *models.User {
	return m.userRepository.UpdateUser(user)
}

func (m *UserApplication) DeleteUser(id int) bool {
	return m.userRepository.DeleteUser(id)
}

func (m *UserApplication) GetUsers() []*models.User {
	return m.userRepository.GetUsers()
}

func (m *UserApplication) GetFollowers() []*models.User {
	return m.userRepository.GetFollowers()
}

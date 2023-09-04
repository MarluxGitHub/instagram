package infrastructure

import (
	"github.com/MarluxGitHub/instagram/pkg/domain/user/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (m *UserRepository) GetUser(id int) *models.User {
	user := models.User{}
	m.DB.First(&user, id)
	return &user
}

func (m *UserRepository) CreateUser(user *models.User) *models.User {
	m.DB.Create(&user)
	return user
}

func (m *UserRepository) UpdateUser(user *models.User) *models.User {
	m.DB.Save(&user)
	return user
}

func (m *UserRepository) DeleteUser(id int) bool {
	m.DB.Delete(&models.User{}, id)
	return true
}

func (m *UserRepository) GetUsers() []*models.User {
	users := []*models.User{}
	m.DB.Find(&users)
	return users
}

func (m *UserRepository) GetFollowers() []*models.User {
	// TODO: implement
	users := []*models.User{}
	m.DB.Find(&users)
	return users
}
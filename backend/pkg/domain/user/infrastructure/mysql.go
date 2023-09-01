package infrastructure

import (
	"github.com/MarluxGitHub/instagram/pkg/domain/user/domain/models"
	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB
}

func (m *MySQL) GetUser(id int) *models.User {
	user := models.User{}
	m.DB.First(&user, id)
	return &user
}

func (m *MySQL) CreateUser(user *models.User) *models.User {
	m.DB.Create(&user)
	return user
}

func (m *MySQL) UpdateUser(user *models.User) *models.User {
	m.DB.Save(&user)
	return user
}

func (m *MySQL) DeleteUser(id int) bool {
	m.DB.Delete(&models.User{}, id)
	return true
}

func (m *MySQL) GetUsers() []*models.User {
	users := []*models.User{}
	m.DB.Find(&users)
	return users
}

func (m *MySQL) GetFollowers() []*models.User {
	// TODO: implement
	users := []*models.User{}
	m.DB.Find(&users)
	return users
}
package migration

import (
	"github.com/MarluxGitHub/instagram/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB ) {
	// Migrate models
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})

}
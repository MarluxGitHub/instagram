package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID 	 uint   `gorm:"primaryKey"`

	Username string `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName string `gorm:"not null"`

	Email string `gorm:"unique;not null"`

	Password string `gorm:"not null"`

	Following []User `gorm:"many2many:following;foreignKey:ID;joinForeignKey:FollowingID;References:ID;JoinReferences:FollowedID"`
}
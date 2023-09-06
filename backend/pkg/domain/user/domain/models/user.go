package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID 	 uint   `gorm:"primaryKey;autoIncrement"`

	Username string `gorm:"unique;not null" json:"username"`
	FirstName string `gorm:"not null" json:"firstName"`
	LastName string `gorm:"not null" json:"lastName"`

	Email string `gorm:"unique;not null" json:"email"`

	Password string `gorm:"not null" json:"password"`

	Following []User `gorm:"many2many:following;foreignKey:ID;joinForeignKey:FollowingID;References:ID;JoinReferences:FollowedID"`
}
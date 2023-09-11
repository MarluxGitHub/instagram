package models

import "gorm.io/gorm"

type Provider struct {
	gorm.Model
	Name string	`json:"name" form:"name"`
	Description string	`json:"description,omitempty" form:"description,omitempty"`
	URI string	`json:"uri" form:"uri"`
}
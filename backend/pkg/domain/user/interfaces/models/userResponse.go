package models

import (
	domainmodels "github.com/MarluxGitHub/instagram/pkg/domain/user/domain/models"
)

type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`

	Followers []UserResponse `json:"followers"`
	Followed  []UserResponse `json:"followed"`
}

func NewUserResponse(user *domainmodels.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
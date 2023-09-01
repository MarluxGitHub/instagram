package application

import "github.com/MarluxGitHub/instagram/pkg/domain/user/domain/services"

type UserApplication struct {
	userService *services.UserService
}
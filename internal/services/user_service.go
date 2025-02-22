package services

import (
	"github.com/v4rakorn1998/go-template/internal/models"
	"github.com/v4rakorn1998/go-template/internal/repositories"
)

func GetUser(req models.UserRequest) ([]models.UserResponse, error) {
	return repositories.GetAllUsers(req)
}

func Register(req models.Register) (bool, error) {
	return repositories.Register(req)
}
func GetUserByUsername(username string) (*models.User, error) {
	return repositories.GetUserByUsername(username)
}

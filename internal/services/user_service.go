package services

import (
	"github.com/v4rakorn1998/go-template/internal/models"
	"github.com/v4rakorn1998/go-template/internal/repositories"
)

func GetUser() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func Register(req models.User) (bool, error) {
	return repositories.Register(req)
}
func GetUserByUsername(username string) (*models.User, error) {
	return repositories.GetUserByUsername(username)
}

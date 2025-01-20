package services

import (
	"github.com/v4rakorn1998/go-template/internal/models"
	"github.com/v4rakorn1998/go-template/internal/repositories"
)

func GetUser() ([]models.User, error) {
	return repositories.GetAllUsers()
}

package services

import (
	"github.com/v4rakorn1998/go-template/internal/models"
	"github.com/v4rakorn1998/go-template/internal/repositories"
)

func GetUser(req models.UserSearchRequest) ([]models.UserResponse, error) {
	return repositories.GetAllUsers(req)
}

func Register(req models.Register) (bool, error) {
	return repositories.Register(req)
}
func GetUserByUsername(username string) (*models.User, error) {
	return repositories.GetUserByUsername(username)
}
func CreateUser(req models.UserRequest) (bool, error) {
	return repositories.CreateUser(req)
}
func UpdateUser(id int, req models.UpdateUserRequest) (bool, error) {
	return repositories.UpdateUser(id, req)
}
func DeleteUser(id int, actionBy string) (bool, error) {
	return repositories.DeleteUser(id, actionBy)
}

func ChangePassword(id int, req models.ChangePasswordRequest) (bool, error) {
	return repositories.ChangePassword(id, req)
}

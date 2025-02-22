package services

import (
	"github.com/v4rakorn1998/go-template/internal/models"
	"github.com/v4rakorn1998/go-template/internal/repositories"
)

func GetCustomersList(request models.CustomerRequest) ([]models.CustomerResponse, error) {
	return repositories.GetCustomersList(request)
}

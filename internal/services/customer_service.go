package services

import (
	"github.com/v4rakorn1998/go-template/internal/models"
	"github.com/v4rakorn1998/go-template/internal/repositories"
)

func GetCustomersList(request models.CustomerSearchRequest) ([]models.CustomerResponse, error) {
	return repositories.GetCustomersList(request)
}
func CreateCustomer(request models.CustomerRequest) (bool, error) {
	return repositories.CreateCustomer(request)
}
func UpdateCustomer(id int, request models.CustomerRequest) (bool, error) {
	return repositories.UpdateCustomer(id, request)
}
func DeleteCustomer(id int, actionBy string) (bool, error) {
	return repositories.DeleteCustomer(id, actionBy)
}

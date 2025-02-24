package handlers

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/v4rakorn1998/go-template/internal/models"
	"github.com/v4rakorn1998/go-template/internal/services"
)

func GetCustomersList(c *fiber.Ctx) error {

	request := new(models.CustomerSearchRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// ตรวจสอบข้อมูลด้วย Validator
	if err := models.Validate.Struct(request); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("Failed on '%s' condition", err.Tag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	customer, err := services.GetCustomersList(*request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customer)
}

func CreateCustomer(c *fiber.Ctx) error {

	request := new(models.CustomerRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// ตรวจสอบข้อมูลด้วย Validator
	if err := models.Validate.Struct(request); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("Failed on '%s' condition", err.Tag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	actionBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	} else {
		request.ActionBy = actionBy
	}

	customer, err := services.CreateCustomer(*request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {

	idParam := c.Params("id")

	// แปลงค่า ID เป็น int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid customer ID"})
	}

	request := new(models.CustomerRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// ตรวจสอบข้อมูลด้วย Validator
	if err := models.Validate.Struct(request); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("Failed on '%s' condition", err.Tag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	actionBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	} else {
		request.ActionBy = actionBy
	}

	customer, err := services.UpdateCustomer(id, *request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customer)
}

func DeleteCustomer(c *fiber.Ctx) error {

	idParam := c.Params("id")

	// แปลงค่า ID เป็น int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid customer ID"})
	}

	actionBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	}
	customer, err := services.DeleteCustomer(id, actionBy)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customer)
}

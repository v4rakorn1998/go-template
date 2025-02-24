package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/v4rakorn1998/go-template/internal/models"
	"github.com/v4rakorn1998/go-template/internal/services"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c *fiber.Ctx) error {
	request := new(models.UserSearchRequest)
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

	users, err := services.GetUser(*request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

// ฟังก์ชันสำหรับสมัครสมาชิก
func CreateUser(c *fiber.Ctx) error {
	request := new(models.UserRequest)
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

	currentTime := time.Now()

	// ดึงเดือนและปี
	month := currentTime.Format("Jan") // รูปแบบเดือนย่อ (Jan, Feb, Mar ฯลฯ)
	year := currentTime.Year()         // ปี (เช่น 2024)

	// รวมเดือนและปีเป็น string
	request.Password = fmt.Sprintf("%s@%d", month, year)

	// เข้ารหัสรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	request.Password = string(hashedPassword)

	// เรียกใช้ service Register และรับค่าผลลัพธ์พร้อมข้อผิดพลาด
	res, err := services.CreateUser(*request)
	if err != nil {
		// ถ้ามีข้อผิดพลาดจากการสมัครสมาชิก, ส่งข้อความผิดพลาด
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// ถ้าสำเร็จ, ส่ง response กลับไป
	return c.Status(fiber.StatusOK).JSON(res)

}

func UpdateUser(c *fiber.Ctx) error {

	idParam := c.Params("id")

	// แปลงค่า ID เป็น int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid customer ID"})
	}

	request := new(models.UpdateUserRequest)
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

	customer, err := services.UpdateUser(id, *request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customer)
}

func DeleteUser(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid customer ID"})
	}

	actionBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	}
	customer, err := services.DeleteUser(id, actionBy)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customer)
}

// ฟังก์ชันสำหรับสมัครสมาชิก
func ChangePassword(c *fiber.Ctx) error {

	idParam := c.Params("id")
	// แปลงค่า ID เป็น int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid customer ID"})
	}

	request := new(models.ChangePasswordRequest)
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
	}
	request.ActionBy = actionBy

	user, err := services.GetUserByUsername(request.ActionBy)
	if err != nil {
		// ถ้าผู้ใช้ไม่พบ หรือมีข้อผิดพลาดจากการค้นหา, ส่ง error
		return c.Status(fiber.StatusUnauthorized).SendString("ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง")
	}

	// เปรียบเทียบรหัสผ่านที่ผู้ใช้กรอกกับรหัสผ่านที่เก็บไว้ในฐานข้อมูล
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		// ถ้ารหัสผ่านไม่ถูกต้อง, ส่ง error
		return c.Status(fiber.StatusUnauthorized).SendString("รหัสผ่านเก่าไม่ถูกต้อง")
	}

	// เข้ารหัสรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	request.Password = string(hashedPassword)

	// เรียกใช้ service Register และรับค่าผลลัพธ์พร้อมข้อผิดพลาด
	res, err := services.ChangePassword(id, *request)
	if err != nil {
		// ถ้ามีข้อผิดพลาดจากการสมัครสมาชิก, ส่งข้อความผิดพลาด
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// ถ้าสำเร็จ, ส่ง response กลับไป
	return c.Status(fiber.StatusOK).JSON(res)

}

// ฟังก์ชันสำหรับสมัครสมาชิก
func ResetPassword(c *fiber.Ctx) error {

	idParam := c.Params("id")
	// แปลงค่า ID เป็น int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid customer ID"})
	}
	request := new(models.ChangePasswordRequest)

	actionBy, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	}
	request.ActionBy = actionBy

	currentTime := time.Now()
	// ดึงเดือนและปี
	month := currentTime.Format("Jan") // รูปแบบเดือนย่อ (Jan, Feb, Mar ฯลฯ)
	year := currentTime.Year()         // ปี (เช่น 2024)

	// รวมเดือนและปีเป็น string
	request.Password = fmt.Sprintf("%s@%d", month, year)

	// เข้ารหัสรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	request.Password = string(hashedPassword)

	// เรียกใช้ service Register และรับค่าผลลัพธ์พร้อมข้อผิดพลาด
	res, err := services.ChangePassword(id, *request)
	if err != nil {
		// ถ้ามีข้อผิดพลาดจากการสมัครสมาชิก, ส่งข้อความผิดพลาด
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// ถ้าสำเร็จ, ส่ง response กลับไป
	return c.Status(fiber.StatusOK).JSON(res)

}

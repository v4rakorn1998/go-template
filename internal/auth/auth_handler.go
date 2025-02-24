package auth

import (
	"fmt"
	"time"

	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/v4rakorn1998/go-template/config"
	"github.com/v4rakorn1998/go-template/internal/models"
	"github.com/v4rakorn1998/go-template/internal/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	request := new(models.Auth)
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

	user, err := services.GetUserByUsername(request.Username)
	if err != nil {
		// ถ้าผู้ใช้ไม่พบ หรือมีข้อผิดพลาดจากการค้นหา, ส่ง error
		return c.Status(fiber.StatusUnauthorized).SendString("ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง")
	}

	// เปรียบเทียบรหัสผ่านที่ผู้ใช้กรอกกับรหัสผ่านที่เก็บไว้ในฐานข้อมูล
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		// ถ้ารหัสผ่านไม่ถูกต้อง, ส่ง error
		return c.Status(fiber.StatusUnauthorized).SendString("ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง")
	}

	// สร้าง JWT

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"username": user.Username,
		"roleCode": user.RoleCode,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // หมดอายุใน 24 ชั่วโมง
	}

	// เซ็นต์ JWT ด้วยคีย์ลับ
	tk, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"token": tk, "message": "เข้าสู่ระบบสำเร็จ",
	})
}

// ฟังก์ชันสำหรับสมัครสมาชิก
func Register(c *fiber.Ctx) error {
	user := new(models.Register)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// เข้ารหัสรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	user.Password = string(hashedPassword)

	// เรียกใช้ service Register และรับค่าผลลัพธ์พร้อมข้อผิดพลาด
	res, err := services.Register(*user)
	if err != nil {
		// ถ้ามีข้อผิดพลาดจากการสมัครสมาชิก, ส่งข้อความผิดพลาด
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// ถ้าสำเร็จ, ส่ง response กลับไป
	return c.Status(fiber.StatusOK).JSON(res)
}

// AuthMiddleware ฟังก์ชันสำหรับตรวจสอบ JWT
func AuthMiddleware(c *fiber.Ctx) error {
	// อ่าน token จาก header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing authorization token")
	}

	// ลบ "Bearer " ออกจาก token string
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// ตรวจสอบว่า token ถูกต้องหรือไม่
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบว่า token เป็นประเภทที่เราคาดหวังหรือไม่
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}

		// ส่งกลับคีย์ลับสำหรับการตรวจสอบ
		return []byte(config.JWTSecret), nil
	})

	// หากมีข้อผิดพลาดในการ parse หรือ token ไม่ถูกต้อง
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	}

	// เพิ่มข้อมูลจาก token ลงใน context เพื่อใช้ในส่วนอื่น ๆ ของแอป
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["username"] == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token claims")
	}

	// ตั้งค่า username ใน context เพื่อใช้งานใน controller ต่าง ๆ
	c.Locals("username", claims["username"])
	c.Locals("roleCode", claims["roleCode"])
	// ถ้า token ถูกต้องให้ดำเนินการต่อ
	return c.Next()
}

// AuthMiddleware ฟังก์ชันสำหรับตรวจสอบ JWT
func CheckRoleAdmin(c *fiber.Ctx) error {

	roleCode := c.Locals("roleCode").(string) == "admin"
	if !roleCode {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden: Insufficient permissions")
	}
	// ถ้า token ถูกต้องให้ดำเนินการต่อ
	return c.Next()
}

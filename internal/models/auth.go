package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

// ตัวแปร global สำหรับใช้ validate
var Validate = validator.New()

// Auth ใช้สำหรับข้อมูลผู้ใช้ในการเข้าสู่ระบบ
type Auth struct {
	Username string `json:"username"  validate:"required"` // ชื่อผู้ใช้
	Password string `json:"password"  validate:"required"` // รหัสผ่าน
}

// Claims struct สำหรับ JWT
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

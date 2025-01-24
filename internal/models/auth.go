package models

import (
	"github.com/golang-jwt/jwt/v4"
)

// Auth ใช้สำหรับข้อมูลผู้ใช้ในการเข้าสู่ระบบ
type Auth struct {
	Username string `json:"username"` // ชื่อผู้ใช้
	Password string `json:"password"` // รหัสผ่าน
}

// Claims struct สำหรับ JWT
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

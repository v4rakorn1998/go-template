package config

import (
	"os"

	"github.com/joho/godotenv"
)

var JWTSecret = os.Getenv("SECRET_KEY")

func Load() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file")
	}

	// ดึงค่า SECRET_KEY หลังจากที่โหลด .env แล้ว
	JWTSecret = os.Getenv("SECRET_KEY")
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

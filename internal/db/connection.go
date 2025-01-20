package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // ใช้ไดรเวอร์ PostgreSQL
)

var DB *sql.DB

func ConnectDB() {
	// กำหนด DSN สำหรับ PostgreSQL
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// สร้าง DSN (Data Source Name) สำหรับ PostgreSQL
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", dbUser, dbPassword, dbName, dbHost, dbPort)
	// เปิดการเชื่อมต่อกับฐานข้อมูล PostgreSQL
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// ตรวจสอบการเชื่อมต่อ
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("Database connected successfully")
}

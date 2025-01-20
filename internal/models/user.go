package models

import "time"

// โครงสร้างของ User ที่มีฟิลด์ Username, Password และ CreatedAt
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

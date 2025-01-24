package models

// โครงสร้างของ User ที่มีฟิลด์ Username, Password และ CreatedAt
type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Status      int    `json:"status"`
	CreatedBy   string `json:"created_by"`
	CreatedDate string `json:"created_date"`
	UpdatedBy   string `json:"updated_by"`
	UpdatedDate string `json:"updated_date"`
}

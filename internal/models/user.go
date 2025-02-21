package models

// โครงสร้างของ User ที่มีฟิลด์ Username, Password และ CreatedAt
type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	RoleCode    string `json:"role_code"`
	Status      bool   `json:"status"`
	CreatedBy   string `json:"created_by"`
	CreatedDate string `json:"created_date"`
	UpdatedBy   string `json:"updated_by"`
	UpdatedDate string `json:"updated_date"`
}

type Register struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	RoleCode    string `json:"roleCode"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dob"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	PicUrl      string `json:"picUrl"`
}

package models

// โครงสร้างของ User ที่มีฟิลด์ Username, Password และ CreatedAt
type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	RoleCode    string `json:"roleCode"`
	Status      bool   `json:"status"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
}

type UserSearchRequest struct {
	PageNumber int    `json:"pageNumber" validate:"required,min=1"`
	PageSize   int    `json:"pageSize" validate:"required,min=1"`
	SearchName string `json:"SearchName"`
}

type UserResponse struct {
	RowNumber         int    `json:"rowNumber"`
	UserID            int    `json:"userID"`
	Username          string `json:"username"`
	RoleCode          string `json:"roleCode"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	DateOfBirth       string `json:"dob"`
	Address           string `json:"address"`
	PhoneNumber       string `json:"phoneNumber"`
	Email             string `json:"email"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
	Status            bool   `json:"status"`
	CreatedBy         string `json:"createdBy"`
	CreatedDate       string `json:"createdDate"`
	UpdatedBy         string `json:"updatedBy"`
	UpdatedDate       string `json:"updatedDate"`
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

type UserRequest struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password"`
	RoleCode    string `json:"roleCode" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	DateOfBirth string `json:"dob" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PicUrl      string `json:"picUrl"`
}
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	Password    string `json:"password" validate:"required"`
	ActionBy    string `json:"actionBy"`
}

type UpdateUserRequest struct {
	RoleCode    string `json:"roleCode" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	DateOfBirth string `json:"dob" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Email       string `json:"email" validate:"required"`
	ActionBy    string `json:"actionBy"`
}

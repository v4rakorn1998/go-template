package models

// โครงสร้างของ User ที่มีฟิลด์ Username, Password และ CreatedAt
type CustomerSearchRequest struct {
	PageNumber  int    `json:"pageNumber" validate:"required,min=1"`
	PageSize    int    `json:"pageSize" validate:"required,min=1"`
	SearchName  string `json:"SearchName"`
	SearchTaxID string `json:"SearchTaxID"`
}

type CustomerResponse struct {
	TotalCount  int    `json:"totalCount"`
	RowNumber   int    `json:"rowNumber"`
	CustomerID  int    `json:"customerID"`
	TypeCode    string `json:"typeCode"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	CompanyName string `json:"companyName"`
	TaxID       string `json:"taxID"`
	Status      bool   `json:"status"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
}

type CustomerRequest struct {
	TypeCode    string `json:"typeCode" validate:"required"`
	FullName    string `json:"fullName" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Email       string `json:"email" validate:"required"`
	CompanyName string `json:"companyName"`
	TaxID       string `json:"taxID" validate:"required"`
	ActionBy    string `json:"actionBy"`
}

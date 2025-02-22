package models

// โครงสร้างของ User ที่มีฟิลด์ Username, Password และ CreatedAt
type CustomerRequest struct {
	PageNumber  int    `json:"pageNumber"`
	PageSize    int    `json:"pageSize"`
	SearchName  string `json:"SearchName"`
	SearchTaxID string `json:"SearchTaxID"`
}

type CustomerResponse struct {
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

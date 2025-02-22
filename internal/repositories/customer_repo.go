package repositories

import (
	"log"

	"github.com/v4rakorn1998/go-template/internal/db"
	"github.com/v4rakorn1998/go-template/internal/models"
)

func GetCustomersList(req models.CustomerRequest) ([]models.CustomerResponse, error) {

	sql := `SELECT row_num,customer_id,type_code,full_name,address,phone_number,email,company_name,tax_id,status,created_by,created_date,updated_by,updated_date FROM fn_get_customer_page($1, $2, $3,$4)`

	rows, err := db.DB.Query(sql, req.PageNumber, req.PageSize, req.SearchName, req.SearchTaxID)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()
	var customers []models.CustomerResponse
	for rows.Next() {
		var customer models.CustomerResponse
		if err := rows.Scan(&customer.RowNumber,
			&customer.CustomerID,
			&customer.TypeCode,
			&customer.FullName,
			&customer.Address,
			&customer.PhoneNumber,
			&customer.Email,
			&customer.CompanyName,
			&customer.TaxID,
			&customer.Status,
			&customer.CreatedBy,
			&customer.CreatedDate,
			&customer.UpdatedBy,
			&customer.UpdatedDate,
		); err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

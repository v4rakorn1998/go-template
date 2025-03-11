package repositories

import (
	"github.com/v4rakorn1998/go-template/internal/db"
	"github.com/v4rakorn1998/go-template/internal/models"
)

func GetCustomersList(req models.CustomerSearchRequest) ([]models.CustomerResponse, error) {

	sql := `SELECT total_count,row_num,customer_id,type_code,full_name,address,phone_number,email,company_name,tax_id,status,created_by,created_date,updated_by,updated_date FROM fn_get_customer_page($1, $2, $3,$4)`

	rows, err := db.DB.Query(sql, req.PageNumber, req.PageSize, req.SearchName, req.SearchTaxID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customers []models.CustomerResponse
	for rows.Next() {
		var customer models.CustomerResponse
		if err := rows.Scan(
			&customer.TotalCount,
			&customer.RowNumber,
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

func CreateCustomer(req models.CustomerRequest) (bool, error) {
	// เริ่มต้น Transaction

	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}

	sqlDetail := "INSERT INTO customer (type_code, full_name, address, phone_number, email, company_name, tax_id, created_by, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err = tx.Exec(sqlDetail, req.TypeCode, req.FullName, req.Address, req.PhoneNumber, req.Email, req.CompanyName, req.TaxID, req.ActionBy, req.ActionBy)
	if err != nil {
		// ถ้าเกิดข้อผิดพลาดในการ insert users_detail, Rollback transaction
		tx.Rollback()
		return false, err
	}

	// ถ้าทุกอย่างสำเร็จ, ทำการ commit transaction
	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

func UpdateCustomer(id int, req models.CustomerRequest) (bool, error) {
	// เริ่มต้น Transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}

	sqlDetail := `UPDATE customer SET type_code= $1, full_name= $2, address= $3, phone_number= $4, email= $5, company_name= $6, tax_id= $7, updated_by= $8, updated_date=NOW() WHERE id=$9;`
	_, err = tx.Exec(sqlDetail, req.TypeCode, req.FullName, req.Address, req.PhoneNumber, req.Email, req.CompanyName, req.TaxID, req.ActionBy, id)
	if err != nil {
		// ถ้าเกิดข้อผิดพลาดในการ insert users_detail, Rollback transaction
		tx.Rollback()
		return false, err
	}

	// ถ้าทุกอย่างสำเร็จ, ทำการ commit transaction
	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteCustomer(id int, actionBy string) (bool, error) {
	// เริ่มต้น Transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}

	sqlDetail := `UPDATE customer SET status = false , updated_by= $1, updated_date=NOW() WHERE id=$2;`
	_, err = tx.Exec(sqlDetail, actionBy, id)
	if err != nil {
		// ถ้าเกิดข้อผิดพลาดในการ insert users_detail, Rollback transaction
		tx.Rollback()
		return false, err
	}

	// ถ้าทุกอย่างสำเร็จ, ทำการ commit transaction
	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

package repositories

import (
	"fmt"

	"github.com/v4rakorn1998/go-template/internal/db"
	"github.com/v4rakorn1998/go-template/internal/models"
)

func GetAllUsers(req models.UserSearchRequest) ([]models.UserResponse, error) {

	sql := `SELECT total_count,row_num,user_id,username,role_code,role_name,status,created_by,created_date,updated_by,updated_date,first_name,last_name,date_of_birth,address,phone_number,email,profile_picture_url 
			FROM public.fn_get_users_page($1, $2, $3)`

	rows, err := db.DB.Query(sql, req.PageNumber, req.PageSize, req.SearchName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(
			&user.TotalCount,
			&user.RowNumber,
			&user.UserID,
			&user.Username,
			&user.RoleCode,
			&user.RoleName,
			&user.Status,
			&user.CreatedBy,
			&user.CreatedDate,
			&user.UpdatedBy,
			&user.UpdatedDate,
			&user.FirstName,
			&user.LastName,
			&user.DateOfBirth,
			&user.Address,
			&user.PhoneNumber,
			&user.Email,
			&user.ProfilePictureUrl,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func Register(req models.Register) (bool, error) {
	// เริ่มต้น Transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}

	// SQL query สำหรับการ insert ข้อมูลในตาราง users พร้อม RETURNING id
	sql := "INSERT INTO users (username, password, role_code) VALUES ($1, $2, $3) RETURNING id"
	var userID int

	// ใช้ QueryRow แทน Query เพื่อรับค่าจากการ return ของ PostgreSQL
	err = tx.QueryRow(sql, req.Username, req.Password, req.RoleCode).Scan(&userID)
	if err != nil {
		// ถ้าเกิดข้อผิดพลาดในการ insert user, Rollback transaction
		tx.Rollback()
		return false, err
	}

	// หลังจากได้ userID แล้ว, ทำการ insert ข้อมูลลงใน users_detail
	sqlDetail := "INSERT INTO users_detail (user_id, first_name, last_name, date_of_birth, address, phone_number, email, profile_picture_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = tx.Exec(sqlDetail, userID, req.FirstName, req.LastName, req.DateOfBirth, req.Address, req.PhoneNumber, req.Email, req.PicUrl)
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

func CreateUser(req models.UserRequest) (bool, error) {
	// เริ่มต้น Transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}

	// SQL query สำหรับการ insert ข้อมูลในตาราง users พร้อม RETURNING id
	sql := "INSERT INTO users (username, password, role_code) VALUES ($1, $2, $3) RETURNING id"
	var userID int

	// ใช้ QueryRow แทน Query เพื่อรับค่าจากการ return ของ PostgreSQL
	err = tx.QueryRow(sql, req.Username, req.Password, req.RoleCode).Scan(&userID)
	if err != nil {
		// ถ้าเกิดข้อผิดพลาดในการ insert user, Rollback transaction
		tx.Rollback()
		return false, err
	}

	// หลังจากได้ userID แล้ว, ทำการ insert ข้อมูลลงใน users_detail
	sqlDetail := "INSERT INTO users_detail (user_id, first_name, last_name, date_of_birth, address, phone_number, email, profile_picture_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = tx.Exec(sqlDetail, userID, req.FirstName, req.LastName, req.DateOfBirth, req.Address, req.PhoneNumber, req.Email, req.PicUrl)
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

func UpdateUser(id int, req models.UpdateUserRequest) (bool, error) {
	// เริ่มต้น Transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}

	sql := `SELECT id, username, password, status, role_code FROM users WHERE status = true  AND id = $1`
	row := tx.QueryRow(sql, id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Status, &user.RoleCode); err != nil {
		return false, fmt.Errorf("user with id %d not found", id)
	}

	sqlDetail := `UPDATE users_detail SET first_name = $1 , last_name = $2, date_of_birth = $3, address = $4, phone_number = $5, email = $6, updated_by = $7 , updated_date=NOW() WHERE user_id= $8;`
	_, err = tx.Exec(sqlDetail, req.FirstName, req.LastName, req.DateOfBirth, req.Address, req.PhoneNumber, req.Email, req.ActionBy, id)
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

func DeleteUser(id int, action string) (bool, error) {
	// เริ่มต้น Transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}

	sqlDetail := `UPDATE users SET status = false , updated_by = $1 , updated_date=NOW() WHERE id= $2;`
	_, err = tx.Exec(sqlDetail, action, id)
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

// ฟังก์ชันสำหรับค้นหาผู้ใช้จากชื่อผู้ใช้
func GetUserByUsername(username string) (*models.User, error) {

	sql := `SELECT id, username, password, status, role_code FROM users WHERE status = true  AND username = $1`

	row := db.DB.QueryRow(sql, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Status, &user.RoleCode)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func ChangePassword(id int, req models.ChangePasswordRequest) (bool, error) {
	// เริ่มต้น Transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}

	sqlDetail := `UPDATE users SET password = $1 , updated_by = $2, updated_date=NOW() WHERE id= $3;`
	_, err = tx.Exec(sqlDetail, req.Password, req.ActionBy, id)
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

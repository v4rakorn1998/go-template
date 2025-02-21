package repositories

import (
	"log"

	"github.com/v4rakorn1998/go-template/internal/db"
	"github.com/v4rakorn1998/go-template/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := db.DB.Query("SELECT id, username, password, role_code, status, created_by, created_date, updated_by, updated_date FROM users")
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.RoleCode, &user.Status, &user.CreatedBy, &user.CreatedDate, &user.UpdatedBy, &user.UpdatedDate); err != nil {
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
		log.Println("Error starting transaction:", err)
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
		log.Println("Error executing query:", err)
		return false, err
	}

	// หลังจากได้ userID แล้ว, ทำการ insert ข้อมูลลงใน users_detail
	sqlDetail := "INSERT INTO users_detail (user_id, first_name, last_name, date_of_birth, address, phone_number, email, profile_picture_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = tx.Exec(sqlDetail, userID, req.FirstName, req.LastName, req.DateOfBirth, req.Address, req.PhoneNumber, req.Email, req.PicUrl)
	if err != nil {
		// ถ้าเกิดข้อผิดพลาดในการ insert users_detail, Rollback transaction
		tx.Rollback()
		log.Println("Error inserting into users_detail:", err)
		return false, err
	}

	// ถ้าทุกอย่างสำเร็จ, ทำการ commit transaction
	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction:", err)
		return false, err
	}

	return true, nil
}

// ฟังก์ชันสำหรับค้นหาผู้ใช้จากชื่อผู้ใช้
func GetUserByUsername(username string) (*models.User, error) {

	sql := `SELECT id, username, password, status FROM users WHERE status = true  AND username = $1`

	row := db.DB.QueryRow(sql, username)

	var user models.User

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Status)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

package repositories

import (
	"log"

	"github.com/v4rakorn1998/go-template/internal/db"
	"github.com/v4rakorn1998/go-template/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := db.DB.Query("SELECT id, username, password, status, created_by, created_date, updated_by, updated_date FROM authentication")
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Status, &user.CreatedBy, &user.CreatedDate, &user.UpdatedBy, &user.UpdatedDate); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func Register(req models.User) (bool, error) {

	sql := "INSERT INTO public.authentication(username, password) VALUES ($1, $2)"

	_, err := db.DB.Query(sql, req.Username, req.Password)
	if err != nil {
		log.Println("Error executing query:", err)
		return false, err
	}

	return true, nil
}

// ฟังก์ชันสำหรับค้นหาผู้ใช้จากชื่อผู้ใช้
func GetUserByUsername(username string) (*models.User, error) {

	sql := `SELECT id, username, password, status FROM public.authentication WHERE status = 1  AND username = $1`

	row := db.DB.QueryRow(sql, username)

	var user models.User

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Status)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

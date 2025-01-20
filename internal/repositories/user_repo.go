package repositories

import (
	"log"

	"github.com/v4rakorn1998/go-template/internal/db"
	"github.com/v4rakorn1998/go-template/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := db.DB.Query("SELECT id , username , email , password , created_at  FROM public.users")
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

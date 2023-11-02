package service

import (
	"database/sql"
	"errors"

	"github.com/Live-Quiz-Project/Backend/internal/db"
	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	"golang.org/x/crypto/bcrypt"
)

func LogIn(email, password string) (*user.User, error) {
	db := db.DB
	var user user.User
	err := db.QueryRow(`SELECT id, email, "password", "name", created_date, account_status FROM "user" WHERE email = $1`, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedDate, &user.AccountStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}

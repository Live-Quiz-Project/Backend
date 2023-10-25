package middleware

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Live-Quiz-Project/Backend/internal/db"
	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(newUser *user.User) error {
	db := db.DB
	newUser.ID = uuid.New().String()

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPassword)

	// Check if email already exists
	var existingUser user.User
	err = db.QueryRow(`SELECT id FROM "user" WHERE email = $1`, newUser.Email).Scan(&existingUser.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Email does not exist, proceed to create new user
		} else {
			log.Println("Error checking for existing user:", err)
			return err
		}
	} else {
		return errors.New("email already taken")
	}

	// Save the user to the database
	_, err = db.Exec(`INSERT INTO "user" (id, email, "password", profile_name, profile_pic, created_date, account_status) VALUES ($1, $2, $3, $4, $5, NOW(), TRUE)`,
		newUser.ID, newUser.Email, newUser.Password, newUser.ProfileName, newUser.ProfilePic)
	if err != nil {
		return err
	}

	return nil
}
func Login(email, password string) (*user.User, error) {
	db := db.DB
	var user user.User
	err := db.QueryRow(`SELECT id, email, "password", profile_name, profile_pic, created_date, account_status FROM "user" WHERE email = $1`, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.ProfileName, &user.ProfilePic, &user.CreatedDate, &user.AccountStatus)
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

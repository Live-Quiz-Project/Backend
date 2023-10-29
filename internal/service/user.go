package service

import (
	"database/sql"
	"errors"
	"log"

	// "strconv"

	"github.com/Live-Quiz-Project/Backend/internal/db"
	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(newUser *user.User) error {
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
	_, err = db.Exec(`INSERT INTO "user" (id, email, "password", "name", created_date, account_status) VALUES ($1, $2, $3, $4, NOW(), TRUE)`,
		newUser.ID, newUser.Email, newUser.Password, newUser.Name)
	if err != nil {
		log.Println("Error saving user to the database:", err)
		return err
	}

	log.Println("User created successfully:", newUser.ID)
	return nil
}

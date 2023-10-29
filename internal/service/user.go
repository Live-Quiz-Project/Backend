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
	_, err = db.Exec(`INSERT INTO "user" (id, email, "password", "name", created_date, account_status) VALUES ($1, $2, $3, $4, NOW(), FALSE)`,
		newUser.ID, newUser.Email, newUser.Password, newUser.Name)
	if err != nil {
		log.Println("Error saving user to the database:", err)
		return err
	}

	log.Println("User created successfully:", newUser.ID)
	return nil
}

func GetUserByID(userID string) (*user.User, error) {
	var user user.User
	err := db.DB.QueryRow(`SELECT id, email, "password", "name", created_date, account_status FROM "user" WHERE id = $1`, userID).
		Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedDate, &user.AccountStatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		log.Println("Error retrieving user from database:", err)
		return nil, err
	}

	return &user, nil
}

func DeleteUser(userID string) error {
	db := db.DB

	_, err := db.Exec(`DELETE FROM "user" WHERE id = $1`, userID)
	if err != nil {
		log.Println("Error deleting user from the database:", err)
		return err
	}

	log.Println("User deleted successfully:", userID)
	return nil
}

func UpdateUser(userID string, updatedName string, updatedImage string) error {
	db := db.DB

	_, err := db.Exec(`UPDATE "user" SET "name" = $1, "image" = $2 WHERE id = $3`,
		updatedName, updatedImage, userID)
	if err != nil {
		log.Println("Error updating user in the database:", err)
		return err
	}

	log.Println("User updated successfully:", userID)
	return nil
}

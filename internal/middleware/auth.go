package middleware

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/Live-Quiz-Project/Backend/internal/db"
	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(newUser *user.User) error {
	db := db.DB

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPassword)

	// Check if email already exists
	var existingUserID string
	err = db.QueryRow(`SELECT id FROM "user" WHERE email = $1`, newUser.Email).Scan(&existingUserID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println("Error checking for existing user:", err)
			return err
		}
	} else {
		return errors.New("email already taken")
	}

	// Retrieve the maximum ID
	var maxID string
	err = db.QueryRow(`SELECT MAX(id) FROM "user"`).Scan(&maxID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error retrieving max user ID:", err)
		return err
	}

	// Convert maxID to integer and increment by 1
	nextID, err := strconv.Atoi(maxID)
	if err != nil && maxID != "" { // If maxID is not an empty string and not a valid integer
		log.Println("Error converting max user ID to integer:", err)
		return err
	}
	nextID++

	// Convert nextID back to string
	newUser.ID = strconv.Itoa(nextID)

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

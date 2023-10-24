package middleware

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Live-Quiz-Project/Backend/internal/db"
	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService() *UserService {
	return &UserService{
		DB: db.DB,
	}
}

func (us *UserService) Register(newUser *user.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPassword)

	// Check if username or email already exists
	var existingUser user.User
	err = us.DB.QueryRow("SELECT id FROM user WHERE username = $1 OR email = $2", newUser.Username, newUser.Email).Scan(&existingUser.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		} else {
			log.Println("Error checking for existing user:", err)
			return err
		}
	} else {
		return errors.New("username or email already taken")
	}

	// Save the user to the database
	_, err = us.DB.Exec("INSERT INTO user (username, password, email, profile_name, profile_pic, created_date, account_status) VALUES (?, ?, ?, ?, ?, NOW(), 'Active')",
		newUser.Username, newUser.Password, newUser.Email, newUser.ProfileName, newUser.ProfilePic)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) Login(username, password string) (*user.User, error) {
	db := db.DB
	var user user.User
	err := db.QueryRow("SELECT id, username, password, email, profile_name, profile_pic, created_date, account_status FROM user WHERE username = ?", username).Scan(
		&user.ID, &user.Username, &user.Password, &user.Email, &user.ProfileName, &user.ProfilePic, &user.CreatedDate, &user.AccountStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return &user, nil
}

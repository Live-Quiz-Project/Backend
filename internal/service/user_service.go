package service

import (
	"errors"

	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	// Add a database connection or any other dependencies here
}

func NewUserService() *UserService {
	return &UserService{
		// Initialize dependencies here
	}
}

func (us *UserService) Register(user *user.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Implement logic to save the user to the database
	// Make sure to check if the username or email already exists

	return nil
}

func (us *UserService) Login(username, password string) (*user.User, error) {
	// Implement logic to retrieve the user from the database by username
	var user user.User
	// ... your database retrieval logic here

	// Check if user is found and if the password is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return &user, nil
}

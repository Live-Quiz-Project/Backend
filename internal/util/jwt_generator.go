package util

import (
	"time"

	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("yourSecretKey")

type Claims struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(user *user.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token is valid for 1 day
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

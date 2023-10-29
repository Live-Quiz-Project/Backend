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

type RefreshClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(user *user.User) (accessToken string, refreshToken string, err error) {
	// Generate Access Token
	accessExpirationTime := time.Now().Add(24 * time.Hour) // Token is valid for 1 day
	accessClaims := &Claims{
		UserID:   user.ID,
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Generate Refresh Token
	refreshExpirationTime := time.Now().Add(24 * time.Hour) // Token is valid for 1 days
	refreshClaims := &RefreshClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

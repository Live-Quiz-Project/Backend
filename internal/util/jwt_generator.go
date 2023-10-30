package util

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var jwtSecret = []byte("secretKey")

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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
	accessTokenExpiryHours, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	refreshTokenExpiryHours, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR"))
	// Generate Access Token
	accessExpirationTime := time.Now().Add(time.Duration(accessTokenExpiryHours) * time.Hour)
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
	refreshExpirationTime := time.Now().Add(time.Duration(refreshTokenExpiryHours) * time.Hour)
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

package user

import (
	"net/http"

	us "github.com/Live-Quiz-Project/Backend/internal/middleware"
	"github.com/Live-Quiz-Project/Backend/internal/util"
	"github.com/gin-gonic/gin"
)

func LogIn(c *gin.Context) {
	var credentials struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if password and confirmPassword match
	if credentials.Password != credentials.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password and confirm password do not match"})
		return
	}

	user, err := us.LogIn(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a JWT token
	accessToken, refreshToken, err := util.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          user,
	})
}

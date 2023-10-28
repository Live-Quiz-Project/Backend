package user

import (
	"net/http"

	us "github.com/Live-Quiz-Project/Backend/internal/middleware"
	"github.com/Live-Quiz-Project/Backend/internal/util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := us.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a JWT token
	token, err := util.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

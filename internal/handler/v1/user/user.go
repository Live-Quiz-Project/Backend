package user

import (
	"net/http"

	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	userService "github.com/Live-Quiz-Project/Backend/internal/service"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	user.User
	ConfirmPassword string `json:"confirmPassword"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate that password and confirmPassword are the same
	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	// Create the user
	err := userService.CreateUser(&req.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := userService.GetUserByID(userID)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, user)
}

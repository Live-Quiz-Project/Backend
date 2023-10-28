package user

import (
	"net/http"

	us "github.com/Live-Quiz-Project/Backend/internal/middleware"
	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser user.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := us.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

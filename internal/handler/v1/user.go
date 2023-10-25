package user_handler

import (
	"net/http"

	"github.com/Live-Quiz-Project/Backend/internal/middleware"
	"github.com/Live-Quiz-Project/Backend/internal/model/user"
	"github.com/Live-Quiz-Project/Backend/internal/util"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *middleware.UserService
}

func NewUserHandler(userService *middleware.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var newUser user.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := uh.userService.Register(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func (uh *UserHandler) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := uh.userService.Login(credentials.Email, credentials.Password)
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

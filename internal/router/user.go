package router

import (
	user_handler "github.com/Live-Quiz-Project/Backend/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userHandler *user_handler.UserHandler) {
	v1 := r.Group("/v1")
	authGroup := v1.Group("/users")

	authGroup.POST("/register", userHandler.CreateUser)
	authGroup.POST("/login", userHandler.Login)
}

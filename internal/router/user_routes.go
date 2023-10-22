package router

import (
	"github.com/Live-Quiz-Project/Backend/internal/handler/user_handler"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, userHandler *user_handler.UserHandler) {
	authGroup := r.Group("/auth")

	authGroup.POST("/register", userHandler.Register)
	authGroup.POST("/login", userHandler.Login)

}

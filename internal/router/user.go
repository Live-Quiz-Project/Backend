package router

import (
	userHandler "github.com/Live-Quiz-Project/Backend/internal/handler/v1/user"
	authMiddleware "github.com/Live-Quiz-Project/Backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	userR := v1.Group("/users")

	userR.POST("", userHandler.CreateUser)
	userR.GET("/:id", authMiddleware.AuthMiddleware("secretKey"), userHandler.GetUserByID)
	userR.DELETE("/:id", authMiddleware.AuthMiddleware("secretKey"), userHandler.DeleteUser)
	userR.PUT("/:id", authMiddleware.AuthMiddleware("secretKey"), userHandler.UpdateUser)
}

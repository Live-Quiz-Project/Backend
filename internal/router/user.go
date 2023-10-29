package router

import (
	userHandler "github.com/Live-Quiz-Project/Backend/internal/handler/v1/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	userR := v1.Group("/users")

	userR.POST("", userHandler.CreateUser)
	userR.GET("/:id", userHandler.GetUserByID)
	userR.DELETE("/:id", userHandler.DeleteUser)
	userR.PUT("/:id", userHandler.UpdateUser)
}

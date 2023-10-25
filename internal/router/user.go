package router

import (
	ush "github.com/Live-Quiz-Project/Backend/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	authGroup := v1.Group("/users")

	authGroup.POST("/register", ush.CreateUser)
	authGroup.POST("/login", ush.Login)
}

package router

import (
	ush "github.com/Live-Quiz-Project/Backend/internal/handler/v1/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	authGroup := v1.Group("/users")

	authGroup.POST("/createUser", ush.CreateUser)
	authGroup.GET("/login", ush.Login)
}

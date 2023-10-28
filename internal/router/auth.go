package router

import (
	ush "github.com/Live-Quiz-Project/Backend/internal/handler/v1/user"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	v1.GET("/login", ush.Login)
}

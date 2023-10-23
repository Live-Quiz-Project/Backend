package router

import (
	"time"

	uh "github.com/Live-Quiz-Project/Backend/internal/handler/user_handler"
	wsh "github.com/Live-Quiz-Project/Backend/internal/handler/v1/ws"
	wsm "github.com/Live-Quiz-Project/Backend/internal/model/ws"
	us "github.com/Live-Quiz-Project/Backend/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173" || origin == "http://localhost:5174" || origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	wsHub := wsm.NewHub()
	wsHandler := wsh.NewWSHandler(wsHub)
	userService := us.NewUserService()
	userHandler := uh.NewUserHandler(userService)
	go wsHub.Run()

	LiveQuizSessionRoutes(r, wsHandler)
	AuthRoutes(r, userHandler)
	QuizManagementRoutes(r)
}

func Start(addr string) error {
	return r.Run(addr)
}

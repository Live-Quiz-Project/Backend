package router

import (
	"github.com/Live-Quiz-Project/Backend/internal/handler/v1/quiz"
	"github.com/gin-gonic/gin"
)

func QuizManagementRoutes(r *gin.Engine) {
	qr := r.Group("/quiz-mgmt")

	qr.GET("/",quiz.GetQuizzes)
	// qr := GET("/quiz", )

	// wsr := r.Group("/lqses")
	// wsr.GET("/", wsHandler.GetLiveQuizSessions)
	// wsr.GET("/:id", wsHandler.GetLiveQuizSessions)
	// wsr.GET("/:id/host", wsHandler.GetHost)
	// wsr.GET("/:id/participants", wsHandler.GetParticipants)
	// wsr.POST("/", wsHandler.CreateLiveQuizSession)
	// wsr.DELETE("/", wsHandler.EndLiveQuizSession)

	// wsr.GET("/join/:code", wsHandler.JoinLiveQuizSession)
	// // wsr.GET("/start/:code", wsHandler.StartLiveQuizSession)
}

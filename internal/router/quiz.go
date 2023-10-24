package router

import (
	"github.com/Live-Quiz-Project/Backend/internal/handler/v1/quiz"
	"github.com/gin-gonic/gin"
)

func QuizManagementRoutes(r *gin.Engine) {
	qr := r.Group("/v1")

	qr.POST("/users",quiz.CreateUser)
	qr.GET("/users", quiz.GetUsers)


	qr.GET("/quizzes",quiz.GetQuizzes)
	qr.GET("/quizzes/:id", quiz.GetQuizByID)
	qr.POST("/quizzes",quiz.CreateQuiz)
	// qr.DELETE("/quizzes/:id", quiz.DeleteQuizByID)
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

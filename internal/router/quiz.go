package router

import (
	"github.com/Live-Quiz-Project/Backend/internal/handler/v1/quiz"
	"github.com/gin-gonic/gin"
)

func QuizManagementRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	qr := v1.Group("/quiz-mgmt")

	qr.POST("/users", quiz.CreateUser)
	qr.GET("/users", quiz.GetUsers)

	qr.GET("quizzes", quiz.GetQuizzes)

	qr.GET("quizzes/:id", quiz.GetAllQuizzesByUserID)
	qr.POST("/quizzes", quiz.CreateQuiz)
	qr.DELETE("/quizzes/:id", quiz.DeleteQuizByID)
}

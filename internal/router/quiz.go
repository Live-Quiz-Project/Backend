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


	qr.GET("quizzes", quiz.GetAllQuizzes)
	qr.GET("quizzes/:id", quiz.GetQuizDetailByQuizID) 

	//qr.GET("quizzes/:id", quiz.GetListQuizzesByUserID) // Get All Quiz in Profile
	qr.POST("/quizzes", quiz.CreateQuiz) // Create Quiz with Full Detail (Quiz, Question, AllOption)
	qr.DELETE("/quizzes/:id", quiz.SoftDeleteQuizByID) // Soft Delete Quiz By ID

	qr.GET("questions/:quiz_id", quiz.GetQuestionDetailByQuizID)
	
}

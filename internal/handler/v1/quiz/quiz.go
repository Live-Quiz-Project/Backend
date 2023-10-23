package quiz

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Live-Quiz-Project/Backend/internal/db"

	"github.com/Live-Quiz-Project/Backend/internal/model/quiz"
)

func GetQuizzes(c *gin.Context) {
	db := db.DB
	// Execute a database query to retrieve users
	rows, err := db.Query("SELECT id, user_id, title, description, media, created_date, modified_date FROM quiz")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var quizzes []quiz.Quiz // Assuming you have a User struct defined
	for rows.Next() {
		var quiz quiz.Quiz // Create a User struct to store data for each row
		if err := rows.Scan(&quiz.ID, &quiz.Title, &quiz.Description, &quiz.OwnerID, &quiz.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		quizzes = append(quizzes, quiz)
	}

	c.JSON(http.StatusOK, quizzes) // Gin will marshal the 'users' slice into JSON
}

func CreateQuizByUserID(c *gin.Context) {
	var quiz quiz.Quiz

	if err := c.BindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

}

package quiz

import (
	"log"
	"net/http"
	"time"
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/Live-Quiz-Project/Backend/internal/db"

	"github.com/Live-Quiz-Project/Backend/internal/model/quiz"

	"github.com/google/uuid"
)

func CreateUser(c *gin.Context) {
	db := db.DB
	// Parse JSON data from the request body into a User struct
	var user quiz.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a UUID for the userID
	userID, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
		return
	}
	user.ID = userID.String()

	user.CreatedDate = time.Now().Format(time.RFC3339)

	query := `
		INSERT INTO "user" (id, email, password, profile_name, profile_pic, created_date, account_status, suspension_start_date, suspension_end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = db.Exec(query, user.ID, user.Email, user.Password, user.ProfileName, user.ProfilePic, user.CreatedDate, user.AccountStatus, user.SuspensionStartDate, user.SuspensionEndDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Respond with the created user's ID
	c.JSON(http.StatusCreated, gin.H{"userID": user.ID})
}

func GetUsers(c *gin.Context) {
	db := db.DB
	// Execute a database query to retrieve users
	rows, err := db.Query("SELECT id, email, password, profile_name, profile_pic, created_date, account_status, suspension_start_date, suspension_end_date FROM \"user\"")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []quiz.User // Assuming you have a User struct defined
	for rows.Next() {
		var user quiz.User // Create a User struct to store data for each row
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.ProfileName, &user.ProfilePic, &user.CreatedDate, &user.AccountStatus, &user.SuspensionStartDate, &user.SuspensionEndDate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users) // Gin will marshal the 'users' slice into JSON
}

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
		if err := rows.Scan(&quiz.ID, &quiz.UserID, &quiz.Title, &quiz.Description, &quiz.Media, &quiz.CreatedDate, &quiz.ModifiedDate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		quizzes = append(quizzes, quiz)
	}

	c.JSON(http.StatusOK, quizzes) // Gin will marshal the 'users' slice into JSON
}

func GetQuizByID(c *gin.Context) {
	db := db.DB
	quizID := c.Param("id") // Get the quizID from the URL parameter

	// Execute a database query to retrieve the quiz by its ID
	row := db.QueryRow("SELECT id, user_id, title, description, media, created_date, modified_date FROM quiz WHERE id = $1", quizID)

	var quiz quiz.Quiz // Create a Quiz struct to store the data
	err := row.Scan(&quiz.ID, &quiz.UserID, &quiz.Title, &quiz.Description, &quiz.Media, &quiz.CreatedDate, &quiz.ModifiedDate)
	if err != nil {
			if err == sql.ErrNoRows {
					c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
			} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
	}

	c.JSON(http.StatusOK, quiz) // Return the quiz as JSON
}

func CreateQuiz(c *gin.Context) {
	db := db.DB

	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to begin a database transaction"})
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var response quiz.CreateQuizRequest
	if err := c.ShouldBindJSON(&response); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quizID, err := uuid.NewUUID()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
		return
	}

	currentTime := time.Now().Format(time.RFC3339)

	quizRecord := quiz.Quiz{
		ID:           quizID.String(),
		UserID:       response.UserID,
		Title:        response.Title,
		Description:  response.Description,
		Media:        response.Media,
		CreatedDate:  currentTime,
		ModifiedDate: "",
	}

	_, err = tx.Exec("INSERT INTO quiz (id, user_id, title, description, media, created_date, modified_date) VALUES ($1, $2, $3, $4, $5, $6, $7)", quizRecord.ID, quizRecord.UserID, quizRecord.Title, quizRecord.Description, quizRecord.Media, quizRecord.CreatedDate, quizRecord.ModifiedDate)
	if err != nil {
		tx.Rollback()
		log.Printf("Failed to insert quiz into the database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error At Quiz"})
		return
	}

	for _, question := range response.Questions {

		// Generate UUID for question
		questionID, err := uuid.NewUUID()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
			return
		}

		questionRecord := quiz.Question{
			ID:               questionID.String(),
			QuizID:           quizRecord.ID,
			IsParentQuestion: question.IsParentQuestion,
			QuestionID:       question.QuestionID,
			Type:             question.Type,
			Content:          question.Content,
			Note:             question.Note,
			Media:            question.Media,
			TimeLimit:        question.TimeLimit,
			HaveTimeFactor:   question.HaveTimeFactor,
			TimeFactor:       question.TimeFactor,
			Font:             question.Font,
			SelectedUpTo:     question.SelectedUpTo,
		}

		_, err = tx.Exec("INSERT INTO question (id, quiz_id, is_parent_question, question_id, type, \"order\", content, note, media, time_limit, have_time_factor, time_factor, font, selected_up_to) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)", questionRecord.ID, questionRecord.QuizID, questionRecord.IsParentQuestion, questionRecord.QuestionID, questionRecord.Type, questionRecord.Order, questionRecord.Content, questionRecord.Note, questionRecord.Media, questionRecord.TimeLimit, questionRecord.HaveTimeFactor, questionRecord.TimeFactor, questionRecord.Font, questionRecord.SelectedUpTo)
		if err != nil {
			tx.Rollback()
			log.Printf("Failed to insert question into the database: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error At Question"})
			return
		}

		switch question.Type {
		case "choice":
			for _, option := range question.OptionChoice {
				optionID, err := uuid.NewUUID()
				if err != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
					return
				}

				optionRecord := quiz.OptionChoice{
					ID:         optionID.String(),
					QuestionID: questionRecord.ID,
					Order:      option.Order,
					Content:    option.Content,
					Point:      option.Point,
					Color:      option.Color,
					IsCorrect:  option.IsCorrect,
				}

				_, err = tx.Exec("INSERT INTO option_choice (id, question_id, \"order\", content, point, color, is_correct) VALUES ($1, $2, $3, $4, $5, $6, $7)", optionRecord.ID, optionRecord.QuestionID, optionRecord.Order, optionRecord.Content, optionRecord.Point, optionRecord.Color, optionRecord.IsCorrect)
				if err != nil {
					tx.Rollback()
					log.Printf("Failed to insert option_choice into the database: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error at option choice"})
					return
				}
			}
		case "text":
			for _, option := range question.OptionText {
				optionID, err := uuid.NewUUID()
				if err != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
					return
				}

				optionRecord := quiz.OptionText{
					ID:                optionID.String(),
					QuestionID:        questionRecord.ID,
					Content:           option.Content,
					Point:             option.Point,
					HaveCaseSensitive: option.HaveCaseSensitive,
				}

				_, err = tx.Exec("INSERT INTO option_text (id, question_id, content, point, have_case_sensitive) VALUES ($1, $2, $3, $4, $5)", optionRecord.ID, optionRecord.QuestionID, optionRecord.Content, optionRecord.Point, optionRecord.HaveCaseSensitive)
				if err != nil {
					tx.Rollback()
					log.Printf("Failed to insert option_text into the database: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error at Option Text"})
					return
				}

			}
		// case "matching":
		// 	for _, option := range question.OptionMatching {
		// 		optionID, err := uuid.NewUUID()
		// 		if err != nil {
		//      tx.Rollback()
		// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
		// 			return
		// 		}

		// 		optionRecord := quiz.OptionMatching{
		// 			ID:                optionID.String(),
		// 			QuestionID:        questionRecord.ID,
		// 			Content:           option.Content,
		// 			Point:             option.Point,
		// 			HaveCaseSensitive: option.HaveCaseSensitive,
		// 		}

		// 		_, err = db.Exec("INSERT INTO option_text (id, question_id, content, point, have_case_sensitive) VALUES ($1, $2, $3, $4, $5)", optionRecord.ID, optionRecord.QuestionID, optionRecord.Content, optionRecord.Point, optionRecord.HaveCaseSensitive)
		// 		if err != nil {
		//      tx.Rollback()
		// 			log.Printf("Failed to insert option_text into the database: %v", err)
		// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error at Option Text"})
		// 			return
		// 		}

		// 	}
		// case "pin":
		//     for _, option := range question.Type {
		//         optionPin := option.Pin
		//         // Handle OptionPin
		//     }
		default:
			// Handle an unknown or unsupported question type
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported question type"})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit the database transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"quizID": quizRecord.ID})
}

// func DeleteQuizByID(c *gin.Context) {
// 	db := db.DB

// 	quizID := c.Param("id")
// 	log.Printf("Received quizID: %v", quizID) // Add this line for debugging


// 	// Check if the quiz exists in the database
// 	var exists bool
// 	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM quiz WHERE id = $1)", quizID).Scan(&exists)
// 	if err != nil {
// 			log.Printf("Failed to check if the quiz exists: %v", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if the quiz exists"})
// 			return
// 	}

// 	if !exists {
// 			log.Printf("Quiz not found: %v", quizID)
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
// 			return
// 	}

// 	// If the quiz exists, proceed with the deletion
// 	tx, err := db.Begin()
// 	if err != nil {
// 			log.Printf("Failed to begin a database transaction: %v", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to begin a database transaction"})
// 			return
// 	}

// 	// Delete the quiz
// 	_, err = tx.Exec("DELETE FROM quiz WHERE id = $1", quizID)
// 	if err != nil {
// 			log.Printf("Failed to delete the quiz: %v", err)
// 			tx.Rollback()
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the quiz"})
// 			return
// 	}

// 	// Commit the transaction if everything is successful
// 	err = tx.Commit()
// 	if err != nil {
// 			log.Printf("Failed to commit the database transaction: %v", err)
// 			tx.Rollback()
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit the database transaction"})
// 			return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Quiz deleted successfully"})
// }
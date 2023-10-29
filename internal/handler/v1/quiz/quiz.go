package quiz

import (
	"errors"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"github.com/Live-Quiz-Project/Backend/internal/db"

	"github.com/Live-Quiz-Project/Backend/internal/model/quiz"

	"github.com/google/uuid"
)

func CreateUser(c *gin.Context) {
	gormdb := db.GormDB

	var response struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	// Keep value in Response
	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserID, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
		return
	}

	created_date := time.Now().Format(time.RFC3339)

	quiz_data := quiz.User{
		ID:                  UserID.String(),
		Email:               response.Email,
		Password:            response.Password,
		Name:                response.Name,
		CreatedDate:         created_date,
		AccountStatus:       "true",
		SuspensionStartDate: "",
		SuspensionEndDate:   "",
	}

	result := gormdb.Create(&quiz_data)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while creating the user"})
		return
	}

	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, quiz_data)
}

func GetUsers(c *gin.Context) {
	gormdb := db.GormDB
	var user []quiz.User
	result := gormdb.Find(&user)
	if result.Error != nil {
		log.Printf("Error retrieving users: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while retrieving users"})
		return
	}
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, user)
}

func CreateQuiz(c *gin.Context) {
	gormdb := db.GormDB

	tx := gormdb.Begin()

	// JSON FROM FRONTEND
	var response struct {
		UserID      string `json:"userId"`
		Title       string `json:"title"`
		Description string `json:"description"`
		CoverImage  string `json:"coverImage"`
		Questions   []struct {
			Version          string `json:"version"`
			IsParent         bool   `json:"isParent"`
			ParentID         string `json:"parentID"`
			IsParentRequired bool   `json:"isParentRequired"`
			Type             string `json:"type"`
			Order            int    `json:"order"`
			Content          string `json:"content"`
			Note             string `json:"note"`
			Media            string `json:"media"`
			TimeLimit        int    `json:"timeLimit"`
			HaveTimeFactor   bool   `json:"haveTimeFactor"`
			TimeFactor       int    `json:"timeFactor"`
			FontSize         int    `json:"font"`
			LayoutIdx        int    `json:"layoutIdx"`
			SelectedUpTo     int    `json:"selectedUpTo"`

			OptionChoice []struct {
				Order     int     `json:"order"`
				Content   string  `json:"content"`
				Mark      float64 `json:"mark"`
				Color     string  `json:"color"`
				IsCorrect bool    `json:"isCorrect"`
			} `json:"optionChoice,omitempty"`

			OptionText []struct {
				Order             int     `json:"order"`
				Content           string  `json:"content"`
				Mark              float64 `json:"mark"`
				HaveCaseSensitive bool    `json:"haveCaseSensitive"`
			} `json:"optionText,omitempty"`
		} `json:"questions"`
	}

	// Keep value in Response
	if err := c.ShouldBindJSON(&response); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a UUID for the userID
	quizID, err := uuid.NewUUID()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
		return
	}

	currentTime := time.Now().Format(time.RFC3339)

	quizData := quiz.Quiz{
		ID:           quizID.String(),
		CreatorID:    response.UserID,
		Version:      currentTime,
		Title:        response.Title,
		Description:  response.Description,
		CoverImage:   response.CoverImage,
		CreatedDate:  currentTime,
		ModifiedDate: "",
		IsDeleted:    "false",
	}

	// Create the new user in the database

	if err := tx.Create(&quizData).Error; err != nil {
		tx.Rollback()
		log.Printf("Error creating quiz: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while creating the quiz"})
		return
	}

	for _, question := range response.Questions {

		questionID, err := uuid.NewUUID()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
			return
		}

		questionData := quiz.Question{
			ID:               questionID.String(),
			QuizID:           quizID.String(),
			Version:          currentTime,
			IsParent:         question.IsParent,
			ParentID:         question.ParentID,
			IsParentRequired: question.IsParentRequired,
			Type:             question.Type,
			Order:            question.Order,
			Content:          question.Content,
			Note:             question.Note,
			Media:            question.Media,
			TimeLimit:        question.TimeLimit,
			HaveTimeFactor:   question.HaveTimeFactor,
			TimeFactor:       question.TimeFactor,
			FontSize:         question.FontSize,
			LayoutIdx:        question.LayoutIdx,
			SelectedUpTo:     question.SelectedUpTo,
		}

		if err := tx.Create(&questionData).Error; err != nil {
			tx.Rollback()
			log.Printf("Error creating question: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while creating the question"})
			return
		}

		// ----- Option ----- //
		for _, optionChoice := range question.OptionChoice {

			optionChoiceID, err := uuid.NewUUID()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
				return
			}

			optionChoiceData := quiz.OptionChoice{
				ID:         optionChoiceID.String(),
				QuestionID: questionID.String(),
				Version:    currentTime,
				Order:      optionChoice.Order,
				Content:    optionChoice.Content,
				Mark:       optionChoice.Mark,
				Color:      optionChoice.Color,
				IsCorrect:  optionChoice.IsCorrect,
			}

			if err := tx.Create(&optionChoiceData).Error; err != nil {
				tx.Rollback()
				log.Printf("Error creating option_choice: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while creating the option_choice"})
				return
			}
		}

		for _, optionText := range question.OptionText {

			optionTextID, err := uuid.NewUUID()
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
				return
			}

			optionTextData := quiz.OptionText{
				ID:                optionTextID.String(),
				QuestionID:        questionID.String(),
				Version:           currentTime,
				Order:             optionText.Order,
				Content:           optionText.Content,
				Mark:              optionText.Mark,
				HaveCaseSensitive: optionText.HaveCaseSensitive,
			}

			if err := tx.Create(&optionTextData).Error; err != nil {
				tx.Rollback()
				log.Printf("Error creating option_text: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while creating the option_text"})
				return
			}
		}
	}

	tx.Commit()

	c.JSON(http.StatusCreated, quizData)
}

func GetQuizzes(c *gin.Context) {
	gormdb := db.GormDB

	var quiz []quiz.Quiz
	result := gormdb.Find(&quiz)

	if result.Error != nil {
		log.Printf("Error retrieving users: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while retrieving quiz"})
		return
	}
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, quiz)
}

func GetAllQuizzesByUserID(c *gin.Context) {
	gormdb := db.GormDB

	userID := c.Param("id")

	var quizzes []quiz.Quiz
	result := gormdb.Find(&quizzes, "creator_id = ?", userID)

	if result.Error != nil {
		log.Printf("Error retrieving users: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while retrieving quiz"})
		return
	}
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, quizzes)
}

func GetQuizByID(c *gin.Context) {
	gormdb := db.GormDB

	userID := c.Param("id")

	var quiz []quiz.Quiz

	result := gormdb.First(&quiz, "creator_id = ?", userID)

	if result.Error != nil {
		log.Printf("Error retrieving users: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while retrieving quiz by id"})
		return
	}
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, quiz)
}

func DeleteQuizByID(c *gin.Context) {
	gormdb := db.GormDB

	quizID := c.Param("id")

	var quiz quiz.Quiz

	// Query the database to find the quiz by its ID
	result := gormdb.First(&quiz, "id = ?", quizID)

	if result.Error != nil {
		// Handle the error if the quiz is not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while fetching the quiz"})
		}
		return
	}

	// Soft delete the quiz by updating the 'is_deleted' field
	result = gormdb.Model(&quiz).Update("is_deleted", true)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while soft deleting the quiz"})
		return
	}

	// Return a success message or appropriate response
	c.JSON(http.StatusOK, gin.H{"message": "Quiz soft deleted successfully"})
}

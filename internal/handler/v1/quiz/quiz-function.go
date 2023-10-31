package quiz

import (
	"github.com/Live-Quiz-Project/Backend/internal/db"

	"github.com/Live-Quiz-Project/Backend/internal/model/quiz"
)


func GetQuizFromDB(quizID string) (res quiz.Quiz, err error) {
	gormdb := db.GormDB

	var response quiz.Quiz
	result := gormdb.Where("id = ?", quizID).First(&response)
	if result.Error != nil {
			return response, result.Error
	}

	return response, nil
}

func GetQuestionFromDB(questionID string) (res quiz.Question, err error) {
	gormdb := db.GormDB

	var response quiz.Question
	result := gormdb.Where("id = ?", questionID).First(&response)
	if result.Error != nil {
			return response, result.Error
	}

	return response, nil
}

func GetOptionChoiceFromDB(optionChoiceID string) (res quiz.OptionChoice, err error) {
	gormdb := db.GormDB

	var response quiz.OptionChoice
	result := gormdb.Where("id = ?", optionChoiceID).First(&response)
	if result.Error != nil {
			return response, result.Error
	}

	return response, nil
}

func GetOptionTextFromDB(optionTextID string) (res quiz.OptionText, err error) {
	gormdb := db.GormDB

	var response quiz.OptionText
	result := gormdb.Where("id = ?", optionTextID).First(&response)
	if result.Error != nil {
			return response, result.Error
	}

	return response, nil
}

func GetOptionMatchingFromDB(optionMatchingID string) (res quiz.OptionMatching, err error) {
	gormdb := db.GormDB

	var response quiz.OptionMatching
	result := gormdb.Where("id = ?", optionMatchingID).First(&response)
	if result.Error != nil {
			return response, result.Error
	}

	return response, nil
}

func GetOptionPoolFromDB(optionPinID string) (res quiz.OptionPin, err error) {
	gormdb := db.GormDB

	var response quiz.OptionPin
	result := gormdb.Where("id = ?", optionPinID).First(&response)
	if result.Error != nil {
			return response, result.Error
	}

	return response, nil
}
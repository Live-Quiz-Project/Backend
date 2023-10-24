package quiz

type CreateQuizRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	OwnerID     string  `json:"owner_id"`
	Questions		[]CreateQuizRequestQuestion `json:"questions"`
}

type CreateQuizRequestQuestion struct {
	QuestionText  string   `json:"question_text"`
	QuestionType  string 	 `json:"question_type"`
	QuestionOrder int    	 `json:"question_order"`
}


// Quiz represents a quiz in your mock data.
type Quiz struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	OwnerID     string  `json:"owner_id"`
	CreatedAt   string  `json:"created_at"`
}

// Question represents a question in your mock data.
type Question struct {
	ID            string   `json:"id"`
	QuizID        string 	 `json:"quiz_id"`
	QuestionText  string   `json:"question_text"`
	QuestionType  string 	 `json:"question_type"`
	QuestionOrder int    	 `json:"question_order"`
}

type Answer struct {
	ID             string   `json:"id"`
	QuestionID     string   `json:"question_id"`
	AnswerType 		 string   `json:"answer_type"`
	Score          float32  `json:"score"`
}

type AnswerOption struct {
	ID          string  `json:"id"`
	AnswerID    string  `json:"answer_id"`
	OptionText  string 	`json:"option_text"`
	OptionOrder int    	`json:"option_order"`
}
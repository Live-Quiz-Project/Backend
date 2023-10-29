package quiz

type User struct {
	ID                  string `json:"id"`
	Password            string `json:"password"`
	Email               string `json:"email"`
	Name                string `json:"name"`
	Image               string `json:"image"`
	CreatedDate         string `json:"createdDate"`
	AccountStatus       string `json:"accountStatus"`
	SuspensionStartDate string `json:"suspensionStartDate"`
	SuspensionEndDate   string `json:"suspensionEndDate"`
}

func (User) TableName() string {
	return "user"
}

// Quiz represents the "quiz" table.
type Quiz struct {
	ID           string `json:"id"`
	CreatorID    string `json:"userId"`
	Version      string `json:"version"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	CoverImage   string `json:"coverImage"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
	IsDeleted    string `json:"isDeleted"`
}

func (Quiz) TableName() string {
	return "quiz"
}

// Question represents the "question" table.
type Question struct {
	ID               string `json:"id"`
	QuizID           string `json:"quizId"`
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
}

func (Question) TableName() string {
	return "question"
}

// OptionChoice represents the "option_choice" table.
type OptionChoice struct {
	ID         string  `json:"id"`
	QuestionID string  `json:"question_id"`
	Version    string  `json:"version"`
	Order      int     `json:"order"`
	Content    string  `json:"content"`
	Mark       float64 `json:"mark"`
	Color      string  `json:"color"`
	IsCorrect  bool    `json:"is_correct"`
}

func (OptionChoice) TableName() string {
	return "option_choice"
}

// OptionText represents the "option_text" table.
type OptionText struct {
	ID                string  `json:"id"`
	QuestionID        string  `json:"question_id"`
	Version           string  `json:"version"`
	Order             int     `json:"order"`
	Content           string  `json:"content"`
	Mark              float64 `json:"mark"`
	HaveCaseSensitive bool    `json:"have_case_sensitive"`
}

func (OptionText) TableName() string {
	return "option_text"
}

// OptionMatching represents the "option_matching" table.
type OptionMatching struct {
	ID         string  `json:"id"`
	QuestionID string  `json:"question_id"`
	PromptID   string  `json:"prompt_id"`
	OptionID   string  `json:"option_id"`
	Mark       float64 `json:"mark"`
}

// OptionMatchingPool represents the "option_matching_pool" table.
type OptionMatchingPool struct {
	ID               string `json:"id"`
	OptionMatchingID string `json:"option_matching_id"`
	Content          string `json:"content"`
	Type             string `json:"type"`
	Order            int    `json:"order"`
}

// OptionPin represents the "option_pin" table.
type OptionPin struct {
	ID         string  `json:"id"`
	QuestionID string  `json:"question_id"`
	XAxis      int     `json:"x_axis"`
	YAxis      int     `json:"y_axis"`
	Mark       float64 `json:"mark"`
}

// LiveQuizSession represents the "live_quiz_session" table.
type LiveQuizSession struct {
	ID        string   `json:"id"`
	QuizID    string   `json:"quiz_id"`
	UserID    string   `json:"user_id"`
	StartedAt string   `json:"started_at"`
	EndedAt   string   `json:"ended_at"`
	IsExempt  []string `json:"is_exempt"`
	Status    string   `json:"status"`
}

// Participant represents the "participant" table.
type Participant struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	LqsID  string  `json:"lqs_id"`
	Name   string  `json:"name"`
	Mark   float64 `json:"mark"`
	Status string  `json:"status"`
}

// ResponseChoice represents the "response_choice" table.
type ResponseChoice struct {
	ID             string `json:"id"`
	ParticipantID  string `json:"participant_id"`
	OptionChoiceID string `json:"option_choice_id"`
}

// ResponseText represents the "response_text" table.
type ResponseText struct {
	ID            string `json:"id"`
	ParticipantID string `json:"participant_id"`
	OptionTextID  string `json:"option_text_id"`
	Content       string `json:"content"`
}

// ResponseMatching represents the "response_matching" table.
type ResponseMatching struct {
	ID               string `json:"id"`
	ParticipantID    string `json:"participant_id"`
	OptionMatchingID string `json:"option_matching_id"`
	PromptID         string `json:"prompt_id"`
	OptionID         string `json:"option_id"`
}

// ResponsePin represents the "response_pin" table.
type ResponsePin struct {
	ID            string `json:"id"`
	ParticipantID string `json:"participant_id"`
	OptionPinID   string `json:"option_pin_id"`
	XAxis         int    `json:"x_axis"`
	YAxis         int    `json:"y_axis"`
}

// Admin represents the "admin" table.
type Admin struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

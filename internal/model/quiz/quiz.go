package quiz

type GetQuiz struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Media        string `json:"media"`
	CreatedDate  string `json:"created_date"`
	ModifiedDate string `json:"modified_date"`
	Questions    []GetQuestion `json:"questions"`
}

type GetQuestion struct {
	ID              string `json:"id"`
	QuizID          string `json:"quiz_id"`
	IsParentQuestion bool   `json:"is_parent_question"`
	QuestionID      string `json:"question_id"`
	Type            string `json:"type"`
	Order           int    `json:"order"`
	Content         string `json:"content"`
	Note            string `json:"note"`
	Media           string `json:"media"`
	TimeLimit       int    `json:"time_limit"`
	HaveTimeFactor  bool   `json:"have_time_factor"`
	TimeFactor      int    `json:"time_factor"`
	Font            int    `json:"font"`
	SelectedUpTo    int    `json:"selected_up_to"`

	OptionChoice  []GetOptionChoice  `json:"option_choice,omitempty"`
	OptionText    []GetOptionText    `json:"option_text,omitempty"`
	//OptionMatching []GetOptionMatching `json:"option_matching,omitempty"`
	//OptionPin     []GetOptionPin     `json:"option_pin,omitempty"`
}

type GetOptionChoice struct {
	ID        string  `json:"id"`
	QuestionID string  `json:"question_id"`
	Order     int     `json:"order"`
	Content   string  `json:"content"`
	Point     float64 `json:"point"`
	Color     string  `json:"color"`
	IsCorrect bool    `json:"is_correct"`
}

type GetOptionText struct {
	ID               string  `json:"id"`
	QuestionID       string  `json:"question_id"`
	Content          string  `json:"content"`
	Point            float64 `json:"point"`
	HaveCaseSensitive bool    `json:"have_case_sensitive"`
}

type CreateQuizRequest struct {
	UserID 			string  `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Media				string  `json:"media"`
	Questions		[]CreateQuestion `json:"questions"`
}

type CreateQuestion struct {
	IsParentQuestion bool  `json:"is_parent_question"`
	QuestionID 			 string `json:"question_id"`
	Type						 string `json:"type"`
	Order						 int		`json:"order"`
	Content					 string `json:"content"`
	Note						 string `json:"note"`
	Media           string `json:"media"`
	TimeLimit       int    `json:"time_limit"`
	HaveTimeFactor  bool   `json:"have_time_factor"`
	TimeFactor      int    `json:"time_factor"`
	Font            int    `json:"font"`
	SelectedUpTo    int    `json:"selected_up_to"`

	OptionChoice  []CreateOptionChoice  `json:"option_choice,omitempty"`
	OptionText    []CreateOptionText    `json:"option_text,omitempty"`
	OptionMatching []CreateOptionMatching `json:"option_matching,omitempty"`
	OptionPin     []CreateOptionPin     `json:"option_pin,omitempty"`
}

type CreateOptionChoice struct {
	Order     int     `json:"order"`
	Content   string  `json:"content"`
	Point     float64 `json:"point"`
	Color     string  `json:"color"`
	IsCorrect bool    `json:"is_correct"`
}

type CreateOptionText struct {
	Content          string  `json:"content"`
	Point            float64 `json:"point"`
	HaveCaseSensitive bool    `json:"have_case_sensitive"`
}

type CreateOptionMatching struct {
	ID        string `json:"id"`
	QuestionID string `json:"question_id"`
	PromptID  string `json:"prompt_id"`
	OptionID  string `json:"option_id"`
	Point     float64 `json:"point"`
	OptionMatchingPool []CreateOptionMatchingPool `json:"option_matching_pool"`
}

type CreateOptionMatchingPool struct {
	Content          string `json:"content"`
	Type            string `json:"type"`
	Order           int    `json:"order"`
}

type CreateOptionPin struct {
	XAxis     int     `json:"x_axis"`
	YAxis     int     `json:"y_axis"`
	Point     float64 `json:"point"`
}

type User struct {
	ID                 string `json:"id"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	ProfileName        string `json:"profile_name"`
	ProfilePic         string `json:"profile_pic"`
	CreatedDate        string `json:"created_date"`
	AccountStatus      bool   `json:"account_status"`
	SuspensionStartDate string `json:"suspension_start_date"`
	SuspensionEndDate   string `json:"suspension_end_date"`
}

// Quiz represents the "quiz" table.
type Quiz struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Media        string `json:"media"`
	CreatedDate  string `json:"created_date"`
	ModifiedDate string `json:"modified_date"`
}

// Question represents the "question" table.
type Question struct {
	ID              string `json:"id"`
	QuizID          string `json:"quiz_id"`
	IsParentQuestion bool   `json:"is_parent_question"`
	QuestionID      string `json:"question_id"`
	Type            string `json:"type"`
	Order           int    `json:"order"`
	Content         string `json:"content"`
	Note            string `json:"note"`
	Media           string `json:"media"`
	TimeLimit       int    `json:"time_limit"`
	HaveTimeFactor  bool   `json:"have_time_factor"`
	TimeFactor      int    `json:"time_factor"`
	Font            int    `json:"font"`
	SelectedUpTo    int    `json:"selected_up_to"`
}

// OptionChoice represents the "option_choice" table.
type OptionChoice struct {
	ID        string  `json:"id"`
	QuestionID string  `json:"question_id"`
	Order     int     `json:"order"`
	Content   string  `json:"content"`
	Point     float64 `json:"point"`
	Color     string  `json:"color"`
	IsCorrect bool    `json:"is_correct"`
}

// OptionText represents the "option_text" table.
type OptionText struct {
	ID               string  `json:"id"`
	QuestionID       string  `json:"question_id"`
	Content          string  `json:"content"`
	Point            float64 `json:"point"`
	HaveCaseSensitive bool    `json:"have_case_sensitive"`
}

// OptionMatching represents the "option_matching" table.
type OptionMatching struct {
	ID        string `json:"id"`
	QuestionID string `json:"question_id"`
	PromptID  string `json:"prompt_id"`
	OptionID  string `json:"option_id"`
	Point     float64 `json:"point"`
}

// OptionMatchingPool represents the "option_matching_pool" table.
type OptionMatchingPool struct {
	ID              string `json:"id"`
	OptionMatchingID string `json:"option_matching_id"`
	Content          string `json:"content"`
	Type            string `json:"type"`
	Order           int    `json:"order"`
}

// OptionPin represents the "option_pin" table.
type OptionPin struct {
	ID        string  `json:"id"`
	QuestionID string  `json:"question_id"`
	XAxis     int     `json:"x_axis"`
	YAxis     int     `json:"y_axis"`
	Point     float64 `json:"point"`
}

// LiveQuizSession represents the "live_quiz_session" table.
type LiveQuizSession struct {
	ID         string   `json:"id"`
	QuizID     string   `json:"quiz_id"`
	UserID     string   `json:"user_id"`
	StartedAt  string   `json:"started_at"`
	EndedAt    string   `json:"ended_at"`
	IsExempt   []string `json:"is_exempt"`
	Status     string   `json:"status"`
}

// Participant represents the "participant" table.
type Participant struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	LqsID  string  `json:"lqs_id"`
	Name   string  `json:"name"`
	Point  float64 `json:"point"`
	Status string  `json:"status"`
}

// ResponseChoice represents the "response_choice" table.
type ResponseChoice struct {
	ID           string `json:"id"`
	ParticipantID string `json:"participant_id"`
	OptionChoiceID string `json:"option_choice_id"`
}

// ResponseText represents the "response_text" table.
type ResponseText struct {
	ID           string `json:"id"`
	ParticipantID string `json:"participant_id"`
	OptionTextID  string `json:"option_text_id"`
	Content      string `json:"content"`
}

// ResponseMatching represents the "response_matching" table.
type ResponseMatching struct {
	ID           string `json:"id"`
	ParticipantID string `json:"participant_id"`
	OptionMatchingID string `json:"option_matching_id"`
	PromptID     string `json:"prompt_id"`
	OptionID     string `json:"option_id"`
}

// ResponsePin represents the "response_pin" table.
type ResponsePin struct {
	ID        string `json:"id"`
	ParticipantID string `json:"participant_id"`
	OptionPinID string `json:"option_pin_id"`
	XAxis     int    `json:"x_axis"`
	YAxis     int    `json:"y_axis"`
}

// Admin represents the "admin" table.
type Admin struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

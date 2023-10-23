package user

type User struct {
	ID                  string `json:"id"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	Email               string `json:"email"`
	ProfileName         string `json:"profileName"`
	ProfilePic          string `json:"profilePic"`
	CreatedDate         string `json:"createdDate"`
	AccountStatus       string `json:"accountStatus"`
	SuspensionStartDate string `json:"suspensionStartDate"`
	SuspensionEndDate   string `json:"suspensionEndDate"`
}

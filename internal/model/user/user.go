package user

type User struct {
	ID                  string
	Email               string
	Password            string
	Name                string
	Image               string
	CreatedDate         string
	AccountStatus       bool
	SuspensionStartDate string
	SuspensionEndDate   string
}

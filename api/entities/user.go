package entities

type User struct {
	UserId        string `json:"userId"`
	Lastname      string `json:"lastname"`
	FirstName     string `json:"firstname"`
	Email         string `json:"email"`
	Password      string `json:",omitempty"`
	MailConfirmed bool   `json:"mailConfirmed"`
}

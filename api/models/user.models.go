package models

type User struct {
	UserId        string `gorm:"column:userId" json:"userId"`
	Lastname      string `gorm:"column:lastname" json:"lastname"`
	FirstName     string `gorm:"column:firstname" json:"firstname"`
	Email         string `gorm:"column:email" json:"email"`
	Password      string `gorm:"column:password" json:",omitempty"`
	MailConfirmed bool   `gorm:"column:mailConfirmed" json:"mailConfirmed"`
}

type UserInformations struct {
	UserId        string     `gorm:"column:userId;primaryKey" json:"userId"`
	Lastname      string     `gorm:"column:lastname" json:"lastname"`
	FirstName     string     `gorm:"column:firstname" json:"firstname"`
	Email         string     `gorm:"column:email" json:"email"`
	MailConfirmed bool       `gorm:"column:mailConfirmed" json:"mailConfirmed"`
	Projects      []*Project `gorm:"ForeignKey:userId"`
}

type UserCredentials struct {
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Token    string `json:"token"`
}

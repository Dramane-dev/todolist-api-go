package models

type Project struct {
	ProjectId          string `gorm:"column:projectId" json:"projectId"`
	ProjectName        string `gorm:"column:name" json:"projectName"`
	ProjectDescription string `gorm:"column:description" json:"projectDescription"`
	UserId             string `gorm:"column:userId" json:"userId"`
}

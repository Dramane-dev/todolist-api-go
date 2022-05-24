package models

type Project struct {
	ProjectId          string        `gorm:"column:projectId;primaryKey" json:"projectId"`
	ProjectName        string        `gorm:"column:name" json:"projectName"`
	ProjectDescription string        `gorm:"column:description" json:"projectDescription"`
	UserId             string        `gorm:"column:userId" json:"userId"`
	Tasks              []*Task       `gorm:"ForeignKey:projectId"`
	Attachments        []*Attachment `gorm:"ForeignKey:projectId"`
}

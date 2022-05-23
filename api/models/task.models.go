package models

type Task struct {
	TaskId      string `gorm:"column:taskId;primaryKey" json:"taskId"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Status      string `gorm:"column:status;default:todo" json:"status"`
	CreatedAt   string `gorm:"column:created_at" json:"createdAt"`
	ProjectId   string `gorm:"column:projectId" json:"projectId"`
}

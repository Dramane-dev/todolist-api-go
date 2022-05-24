package models

type Attachment struct {
	AttachmentId string `gorm:"column:attachmentId;primaryKey" json:"attachmentId"`
	FileName     string `gorm:"column:name" json:"fileName"`
	FileType     string `gorm:"column:type" json:"fileType"`
	FilePath     string `gorm:"column:path" json:"filePath"`
	ProjectId    string `gorm:"column:projectId" json:"projectId"`
}

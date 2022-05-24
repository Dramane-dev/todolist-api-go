package service

import "github.com/Dramane-dev/todolist-api/api/models"

type AttachmentService interface {
	GetAllAttachments() ([]*models.Attachment, error)
	GetAllAttachmentsByProjectId(projectId string) ([]*models.Attachment, error)
	GetAttachmentById(attachmentId string) (*models.Attachment, error)
	UploadAttachment(attachment *models.Attachment) (*models.Attachment, error)
	DeleteAttachment(attachmentId string) (*string, error)
}

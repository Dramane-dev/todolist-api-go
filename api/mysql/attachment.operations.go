package mysql

import (
	"errors"

	"github.com/Dramane-dev/todolist-api/api/models"
	"github.com/google/uuid"
)

func (db *MySQLDatabase) GetAllAttachments() ([]*models.Attachment, error) {
	var attachments []*models.Attachment

	errWhenGettingAttachmentsProject := db.connection.Model(&models.Attachment{}).Find(&attachments).Error

	if errWhenGettingAttachmentsProject != nil {
		return nil, errWhenGettingAttachmentsProject
	}

	return attachments, nil
}

func (db *MySQLDatabase) GetAllAttachmentsByProjectId(projectId string) ([]*models.Attachment, error) {
	var attachments []*models.Attachment

	errWhenGettingAttachmentsProject := db.connection.Model(&models.Attachment{}).Where("projectId = ?", projectId).Find(&attachments).Error

	if errWhenGettingAttachmentsProject != nil {
		return nil, errWhenGettingAttachmentsProject
	}

	return attachments, nil
}

func (db *MySQLDatabase) GetAttachmentById(attachmentId string) (*models.Attachment, error) {
	var attachment *models.Attachment

	errWhenGetAttachmentById := db.connection.Model(&models.Attachment{}).Where("attachmentId = ?", attachmentId).Find(&attachment).Error

	if errWhenGetAttachmentById != nil {
		return nil, errWhenGetAttachmentById
	}

	return attachment, nil
}

func (db *MySQLDatabase) UploadAttachment(attachment *models.Attachment) (*models.Attachment, error) {
	attachment.AttachmentId = "ATC" + uuid.NewString()
	errWhenCreateProject := db.connection.Model(&models.Attachment{}).Create(attachment).Error

	if errWhenCreateProject != nil {
		return nil, errWhenCreateProject
	}

	return db.GetAttachmentById(attachment.AttachmentId)
}

func (db *MySQLDatabase) DeleteAttachment(attachmentId string) (*string, error) {
	attachment, _ := db.GetAttachmentById(attachmentId)

	if !(len(attachment.AttachmentId) > 0) {
		return nil, errors.New("project not found ...❌")
	}

	errWhenDeleteProject := db.connection.Model(&models.Attachment{}).Where("attachmentId = ?", attachmentId).Delete(attachmentId).Error

	if errWhenDeleteProject != nil {
		return nil, errWhenDeleteProject
	}

	var strResponse string = "Attachment deleted successfully ✅"

	return &strResponse, errWhenDeleteProject
}

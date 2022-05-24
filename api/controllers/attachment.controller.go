package controllers

import (
	"fmt"
	"net/http"

	"github.com/Dramane-dev/todolist-api/api/functions"
	"github.com/Dramane-dev/todolist-api/api/models"

	"github.com/gin-gonic/gin"
)

func (attachmentService *AttachmentController) GetAllAttachments(ctx *gin.Context) {
	attachments, errWhenGetAllAttachmentByUserId := attachmentService.database.GetAllAttachments()

	if errWhenGetAllAttachmentByUserId != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenGetAllAttachmentByUserId})
		return
	}

	data := map[string]interface{}{
		"attachments": attachments,
	}

	ctx.JSON(http.StatusOK, data)
}

func (attachmentService *AttachmentController) GetAllAttachmentsByProjectId(ctx *gin.Context) {
	projectId, ok := ctx.Params.Get("projectId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "projectId not provided or incorrect...❌"})
		return
	}

	attachments, errWhenGetAllAttachments := attachmentService.database.GetAllAttachmentsByProjectId(projectId)

	if errWhenGetAllAttachments != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, errWhenGetAllAttachments)
	}

	data := map[string]interface{}{
		"attachments": attachments,
	}

	ctx.JSON(http.StatusOK, data)
}

func (attachmentService *AttachmentController) GetAttachmentById(ctx *gin.Context) {
	projectId, ok := ctx.Params.Get("projectId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "projectId not provided or incorrect...❌"})
		return
	}

	attachment, errWhenGetAttachmentById := attachmentService.database.GetAttachmentById(projectId)

	if errWhenGetAttachmentById != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenGetAttachmentById})
		return
	}

	if len(attachment.ProjectId) > 0 {
		data := map[string]interface{}{
			"attachment": attachment,
		}

		ctx.JSON(http.StatusOK, data)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Attachment not found...❌"})
}

func (attachmentService *AttachmentController) UploadAttachment(ctx *gin.Context) {
	if err := ctx.Request.ParseMultipartForm(500 << 20); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var attachment *models.Attachment = new(models.Attachment)
	projectId, ok := ctx.Params.Get("projectId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "projectId not provided or incorrect"})
		return
	}

	file, errWhenGettingFile := ctx.FormFile("file")

	if errWhenGettingFile != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errWhenGettingFile.Error()})
		return
	}

	attachment.ProjectId = projectId
	attachment.FileName = file.Filename
	attachment.FileType = file.Header.Get("Content-Type")
	attachment.FilePath = "./uploads/" + file.Filename

	errWhenSaveFile := ctx.SaveUploadedFile(file, attachment.FilePath)

	if errWhenSaveFile != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errWhenSaveFile.Error()})
		return
	}

	// result, errWhenCropedImage := functions.CropAttachment(attachment, 250, 250)

	// if errWhenCropedImage != nil {
	// 	fmt.Println(errWhenCropedImage.Error())
	// }

	messageResponse := make(chan string)

	go func() {
		result, errWhenCropedImage := functions.CropAttachment(attachment, 250, 250)

		if errWhenCropedImage != nil {
			fmt.Println(errWhenCropedImage.Error())
		}

		messageResponse <- *result
	}()

	cropedImgResponse := <-messageResponse

	attachment, errWhenUploadAttachment := attachmentService.database.UploadAttachment(attachment)

	if errWhenUploadAttachment != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errWhenUploadAttachment})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": file.Filename + " upload successfully ✅", "cropedImage": cropedImgResponse, "attachment": attachment})
	// ctx.JSON(http.StatusOK, gin.H{"message": file.Filename + " upload successfully ✅"})
}

func (attachmentService *AttachmentController) DeleteAttachment(ctx *gin.Context) {
	attachmentId, ok := ctx.Params.Get("attachmentId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "attachmentId not provided or incorrect...❌"})
		return
	}

	attachment, errWhenDeleteAttachment := attachmentService.database.DeleteAttachment(attachmentId)

	if errWhenDeleteAttachment != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenDeleteAttachment.Error()})
		return
	}

	ctx.JSON(http.StatusOK, attachment)
}

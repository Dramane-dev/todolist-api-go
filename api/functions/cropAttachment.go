package functions

import (
	"image"
	"log"
	"os"
	"strings"

	"github.com/Dramane-dev/todolist-api/api/models"
	"github.com/disintegration/imaging"
)

func CropAttachment(attachment *models.Attachment, width int, height int) (*string, error) {
	var imgCroped image.Image
	fileReaded, errWhenReadFile := os.Open(attachment.FilePath)

	if errWhenReadFile != nil {
		return nil, errWhenReadFile
	}

	switch true {
	case strings.Contains(attachment.FileType, "png"):
		pngImg, _, errWhenDecodeAttachment := image.Decode(fileReaded)

		if errWhenDecodeAttachment != nil {
			return nil, errWhenDecodeAttachment
		}

		defer fileReaded.Close()

		imgCroped = imaging.Resize(pngImg, width, height, imaging.Lanczos)
	case strings.Contains(attachment.FileType, "jpeg"):
		jpegImg, _, errWhenDecodeAttachment := image.Decode(fileReaded)

		if errWhenDecodeAttachment != nil {
			return nil, errWhenDecodeAttachment
		}

		defer fileReaded.Close()

		imgCroped = imaging.Resize(jpegImg, width, height, imaging.Lanczos)
	}

	errWhenSavedImage := imaging.Save(imgCroped, "./uploads/croped_"+attachment.FileName)

	if errWhenSavedImage != nil {
		return nil, errWhenSavedImage
	}

	mailSent, errorWhenSendAnEmail := SendAnEmail("croped_"+attachment.FileName, "./uploads/croped_"+attachment.FileName)

	if errorWhenSendAnEmail != nil {
		return nil, errorWhenSendAnEmail
	}

	log.Println(mailSent)

	responseMessage := "Attachment croped_" + attachment.FileName + " croped successfully ✅"
	return &responseMessage, nil
}

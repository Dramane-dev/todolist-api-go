package functions

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"

	"github.com/Dramane-dev/todolist-api/api/smtpLoginAuth"
)

type Mail struct {
	From       string
	To         []string
	Subject    string
	Body       string
	Attachment Attachment
}

type Attachment struct {
	Name string
	Data []byte
}

func SendAnEmail(fileName string, filePath string) (*string, error) {
	from := os.Getenv("MAIL_SENDER")
	password := os.Getenv("MAIL_PASSWORD")
	to := []string{
		"dramane.kamissoko@estiam.com",
		"kamissokoo.draamane@gmail.com",
	}
	subject := os.Getenv("MAIL_SUBJECT")
	body := "Bonjour,\n vous trouverez votre image en pièce jointe.\n Bonne réception.\n\n Simply Todo"
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")
	data, errorWhenReadFile := ioutil.ReadFile(filePath)

	if errorWhenReadFile != nil {
		return nil, errorWhenReadFile
	}

	attachment := Attachment{
		Name: fileName,
		Data: data,
	}

	request := Mail{
		From:       from,
		To:         to,
		Subject:    subject,
		Body:       body,
		Attachment: attachment,
	}

	mailContent := EncodingFile(filePath, request)
	auth := smtpLoginAuth.LoginAuth(from, password)
	// auth := smtp.PlainAuth("", from, password, smtpHost)
	addr := smtpHost + ":" + smtpPort
	errWhenSendingMail := smtp.SendMail(addr, auth, from, to, mailContent)

	if errWhenSendingMail != nil {
		return nil, errWhenSendingMail
	}

	response := "Mail sent successfully ✅"
	return &response, nil
}

func EncodingFile(filePath string, request Mail) []byte {
	var buff bytes.Buffer

	buff.WriteString(fmt.Sprintf("From: %s\r\n", request.From))
	buff.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(request.To, ";")))
	buff.WriteString(fmt.Sprintf("Subject: %s\r\n", request.Subject))

	boundary := "my-boundary-779"
	buff.WriteString("MIME-Version: 1.0\r\n")
	buff.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n",
		boundary))

	buff.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buff.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buff.WriteString(fmt.Sprintf("\r\n%s", request.Body))

	buff.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buff.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buff.WriteString("Content-Transfer-Encoding: base64\r\n")
	buff.WriteString("Content-Disposition: attachment; filename=" + request.Attachment.Name + "\r\n")
	buff.WriteString("Content-ID: <" + request.Attachment.Name + ">\r\n\r\n")

	convertedFile := make([]byte, base64.StdEncoding.EncodedLen(len(request.Attachment.Data)))
	base64.StdEncoding.Encode(convertedFile, request.Attachment.Data)
	buff.Write(convertedFile)

	// var imgToBase64Encoding string

	// mimetype := http.DetectContentType(request.Attachment.Data)

	// switch mimetype {
	// case "image/jpeg":
	// 	imgToBase64Encoding += "data:image/jpeg;base64,"
	// case "image/png":
	// 	imgToBase64Encoding += "data:image/png;base64,"
	// }

	// imgToBase64Encoding += toBase64(request.Attachment.Data)
	// // fmt.Println(imgToBase64Encoding)
	// // buff.WriteString(imgToBase64Encoding)

	buff.WriteString(fmt.Sprintf("\r\n--%s", boundary))

	buff.WriteString("--")
	return buff.Bytes()
}

// func toBase64(img []byte) string {
// 	return base64.StdEncoding.EncodeToString(img)
// }

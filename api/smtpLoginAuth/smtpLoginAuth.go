package smtpLoginAuth

import (
	"errors"
	"net/smtp"
)

type loginAuth struct {
	Username string
	password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{Username: username, password: password}
}

func (auth *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (auth *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(auth.Username), nil
		case "Password:":
			return []byte(auth.password), nil
		default:
			return nil, errors.New("unkown fromServer")
		}
	}
	return nil, nil
}

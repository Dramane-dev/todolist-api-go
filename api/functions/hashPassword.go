package functions

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err.Error(), nil
	}

	return string(passwordHashed), err
}

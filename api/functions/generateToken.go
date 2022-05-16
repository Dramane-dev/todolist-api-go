package functions

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserId    string `json:"userId"`
	NotBefore int64  `json:"notBefore"`
	jwt.StandardClaims
}

func GenerateNewToken(userId string) (string, error) {
	customClaims := CustomClaims{
		UserId:    userId,
		NotBefore: time.Now().Add(5 * time.Minute).Unix(),
	}
	acessTokenSecret := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, errWhenSignedToken := token.SignedString(acessTokenSecret)

	if errWhenSignedToken != nil {
		return errWhenSignedToken.Error(), nil
	}
	return tokenString, nil
}

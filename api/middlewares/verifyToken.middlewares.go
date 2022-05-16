package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Dramane-dev/todolist-api/api/functions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		acessTokenSecret := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
		header := c.Request.Header["Authorization"]

		if !(len(header) > 0) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided ...❌ Please Signin with your account or create an account."})
			return
		}

		tokenString := strings.Split(header[0], " ")[1]

		token, errWhenParsingJwt := jwt.ParseWithClaims(tokenString, &functions.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method : %v", token.Header[""])
			}
			return acessTokenSecret, nil
		})

		if claims, ok := token.Claims.(*functions.CustomClaims); ok && token.Valid {
			fmt.Println(claims.UserId, claims.NotBefore)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println(errWhenParsingJwt)
			return
		}
	}
}

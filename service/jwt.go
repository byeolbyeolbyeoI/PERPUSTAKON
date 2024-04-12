package service 

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"

	"perpustakaan/models"
)

// consider whether the variable name should be dbUser
func GenerateJWT(user models.User) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
		"username": user.Username,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	/*
	// fuck around n find out
	fmt.Println("signing method : ", token.Method)
	fmt.Println("raw : ", token.Claims)
	fmt.Println("header : ", token.Header) // so this is where header["alg"] comes from
	fmt.Println("signature : ", token.Signature)
	fmt.Println("valid : ", token.Valid)
	*/
	return token
}

func SignToken(token *jwt.Token) (string, error) {
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

package controller

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtSecretKey = []byte("110C9CD469E84C570B64FEB9E18B09C42F38932335294F6E1726D68C2DC6EAC4") //

// Generate JWT token dengan payload (username,id_role,title)
func GenerateToken(username string, IdRole int, title string) (string, time.Time, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // Token berlaku selama 7 hari
	claims := jwt.MapClaims{
		"username": username,
		"id_role":  IdRole,
		"title":    title,
		"exp":      expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecretKey)
	return tokenString, expirationTime, err
}

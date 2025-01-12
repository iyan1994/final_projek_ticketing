package controller

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func ParseTokenMapClaims(tokenString string) (jwt.MapClaims, error) {
	// Parse token menggunakan MapClaims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signing adalah HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Ambil klaim dari token jika valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

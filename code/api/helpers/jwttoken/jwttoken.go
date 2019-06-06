package jwttoken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	expiresIn = 5
	secret    = "SECRET"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type AuthTokenClaim struct {
	jwt.StandardClaims
	User
}

func GenerateJWTToken(user User) (string, error) {
	expiresAt := time.Now().Add(expiresIn * time.Minute).Unix()
	mySigningKey := []byte(secret)

	claims := AuthTokenClaim{
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "test",
		},
		user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserFromJWTToken(tokenString string) (bool, User) {
	var claims AuthTokenClaim
	var user User
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false, user
	}

	user = claims.User

	return token.Valid, user
}

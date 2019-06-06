package hash

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 14
)

func Generate(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), cost)
	return string(bytes), err
}

func VerifyHash(value, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}

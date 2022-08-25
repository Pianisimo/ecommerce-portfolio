package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func IsPasswordValid(hash string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	check := true
	if err != nil {
		check = false
	}

	return check, err
}

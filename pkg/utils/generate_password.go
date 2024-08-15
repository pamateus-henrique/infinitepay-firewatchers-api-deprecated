package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		fmt.Println("Error generating password")
		return err.Error()
	}

	return string(hashedPassword)

}
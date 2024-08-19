package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(password string, hashedPassword string) (error) {
	
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}

	return nil
}
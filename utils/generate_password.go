package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(request string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request), 15)
	return string(hashPassword), err
}

package middleware

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(Set_user_password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Set_user_password), 14)
	return string(bytes), err
}

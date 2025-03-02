package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// generatePasswordHash hashes the user's password using SHA-512, a salt, and pepper.
func GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// VerifyPassword will be used to verify the user submitted password against the stored passwordHash
func VerifyPassword(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

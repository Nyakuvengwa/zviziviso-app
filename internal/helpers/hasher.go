package helpers

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"strings"
)

// generatePasswordHash hashes the user's password using SHA-512, a salt, and pepper.
func GeneratePasswordHash(password string) (string, error) {
	// Generate a salt
	salt, err := generateSalt(16) // 16 bytes salt
	if err != nil {
		return "", fmt.Errorf("error generating salt: %w", err)
	}

	// Get the pepper from environment variables or a secure config file
	pepper, err := getPepper()
	if err != nil {
		return "", fmt.Errorf("error getting pepper: %w", err)
	}

	// Concatenate the password, salt, and pepper
	saltedPassword := fmt.Sprintf("%s%s%s", password, salt, pepper)

	// Hash the salted password using SHA-512
	hashedPassword := sha512.Sum512([]byte(saltedPassword))

	// Base64 encode the salt and the hashed password
	encodedSalt := base64.StdEncoding.EncodeToString([]byte(salt))
	encodedHash := base64.StdEncoding.EncodeToString(hashedPassword[:])

	// Combine the encoded salt and hash, separated by a $ (or another delimiter)
	finalHash := fmt.Sprintf("%s$%s", encodedSalt, encodedHash)

	return finalHash, nil
}

// generateSalt generates a random salt of the specified length.
func generateSalt(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("error generating random salt: %w", err)
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// getPepper retrieves the pepper from a secure source.
// In a real application, this would likely come from environment variables,
// a secure config file, or a key vault.
func getPepper() (string, error) {
	// IMPORTANT: In a real application, DO NOT hardcode the pepper.
	// This is for demonstration purposes ONLY.
	// Use environment variables, a secure config file, or a key vault.

	//Example using Env Variables
	//pepper := os.Getenv("PASSWORD_PEPPER")
	//if pepper == "" {
	//	return "", fmt.Errorf("PASSWORD_PEPPER environment variable is not set")
	//}
	//return pepper, nil

	//For this demo, it's hardcoded
	pepper := "this-is-a-super-secret-pepper-change-this-for-production"
	if strings.TrimSpace(pepper) == "" {
		return "", fmt.Errorf("Pepper is not configured correctly")
	}
	return pepper, nil
}

// VerifyPassword will be used to verify the user submitted password against the stored passwordHash
func VerifyPassword(submittedPassword string, passwordHash string) (bool, error) {
	parts := strings.Split(passwordHash, "$")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid password hash format")
	}
	encodedSalt := parts[0]
	encodedHash := parts[1]

	salt, err := base64.StdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return false, fmt.Errorf("error decoding salt: %w", err)
	}

	pepper, err := getPepper()
	if err != nil {
		return false, fmt.Errorf("error getting pepper: %w", err)
	}

	saltedPassword := fmt.Sprintf("%s%s%s", submittedPassword, salt, pepper)
	hashedPassword := sha512.Sum512([]byte(saltedPassword))
	calculatedHash := base64.StdEncoding.EncodeToString(hashedPassword[:])
	
	return calculatedHash == encodedHash, nil
}

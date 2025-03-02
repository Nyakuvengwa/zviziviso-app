package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"
	repository "zviziviso-app/internal/db"
)

type CreateUserRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

func (app *Application) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusBadRequest, INVALID_REQUEST, "Invalid request body provided.")
		return
	}

	err = validateUser(user)

	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusBadRequest, INVALID_REQUEST, err.Error())
		return
	}

	//
	users, err := app.service.GetUserByEmailOrUsername(ctx, repository.GetUserByEmailOrUsernameParams{Username: user.Username, Email: user.Email})

	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, err.Error())
		return
	}

	if len(users) > 0 {
		for _, dbUser := range users {
			if dbUser.Email == user.Email {
				NewProblemDetailsErrorResponse(w, http.StatusBadRequest, INVALID_REQUEST, "A user with the provided email address is already registered with our service. please provide a alternative email.")
				return
			}

			if dbUser.Username == user.Username {
				NewProblemDetailsErrorResponse(w, http.StatusBadRequest, INVALID_REQUEST, "A user with the provided username is already registered with our service. please provide a alternative username.")
				return
			}
		}
	}

	//TODO: hash password

	passwordHash := generatePasswordHash(user)

	//TODO: save user to database

	//TODO: return user details

	json.NewEncoder(w).Encode(user)
}

func generatePasswordHash(user CreateUserRequest) any {
	panic("unimplemented")
}

func validateUser(user CreateUserRequest) error {
	errorMessages := make([]string, 5)

	if strings.TrimSpace(user.Username) == "" {
		errorMessages = append(errorMessages, "Username cannot be empty")
	}

	if strings.TrimSpace(user.Email) == "" {
		errorMessages = append(errorMessages, "Email cannot be empty")
	} else {
		_, err := mail.ParseAddress(user.Email)
		errorMessages = append(errorMessages, err.Error())
	}

	if strings.TrimSpace(user.Password) == "" {
		errorMessages = append(errorMessages, "Password cannot be empty")
	}

	if strings.TrimSpace(user.FirstName) == "" {
		errorMessages = append(errorMessages, "First name cannot be empty")
	}

	if strings.TrimSpace(user.LastName) == "" {
		errorMessages = append(errorMessages, "Last name cannot be empty")
	}

	// Check if errorMessages has any strings in it and if it has return error
	if len(errorMessages) > 0 {
		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}
	return nil
}

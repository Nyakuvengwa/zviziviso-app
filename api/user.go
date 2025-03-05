package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"
	repository "zviziviso-app/internal/db"
	"zviziviso-app/internal/helpers"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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

	passwordHash, err := helpers.GeneratePasswordHash(user.Password)

	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, "Error saving user details. Contact support.")
		return
	}

	userId, err := app.service.CreateUserDetails(ctx, repository.CreateUserDetailsParams{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: passwordHash,
		FirstName:    pgtype.Text{String: user.FirstName, Valid: true},
		LastName:     pgtype.Text{String: user.LastName, Valid: true},
		Role:         pgtype.Text{String: "User", Valid: true},
	})

	if err != nil {
		log.Printf("DB Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, "Error saving user details. Contact support.")
		return
	}

	result, err := app.service.GetUserSummaryDetails(ctx, userId)

	if err != nil {
		log.Printf("DB Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusInternalServerError, INTERNAL_SERVER_ERROR, "Error saving user details. Contact support.")
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (app *Application) GetUserSummaryDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.PathValue("userId")
	err := uuid.Validate(userId)
	if err != nil {
		NewProblemDetailsErrorResponse(w, http.StatusBadRequest, "Invalid user id provided", "Invalid url path parameter provided.")
		return
	}
	user, err := app.service.GetUserSummaryDetails(ctx, uuid.MustParse(userId))

	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusInternalServerError, "Unhandled server error", err.Error())
		return
	}
	json.NewEncoder(w).Encode(user)
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
		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
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

	var filteredErrorMessages []string
	for _, msg := range errorMessages {
		if msg != "" {
			filteredErrorMessages = append(filteredErrorMessages, msg)
		}
	}
	// Check if errorMessages has any strings in it and if it has return error
	if len(filteredErrorMessages) > 0 {
		return fmt.Errorf(strings.Join(filteredErrorMessages, ", "))
	}
	return nil
}

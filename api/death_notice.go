package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type createDeathNoticeRequest struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Title        string    `json:"title"`
	DateOfDeath  time.Time `json:"date_of_death"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	CauseOfDeath string    `json:"cause_of_death"`
	Obituary     string    `json:"obituary"`
	ImageUrl     string    `json:"image_url"`
}

func (app *Application) CreateNewDeathNotice(w http.ResponseWriter, r *http.Request) {
	var deathNotice createDeathNoticeRequest
	err := json.NewDecoder(r.Body).Decode(&deathNotice)
	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusBadRequest, INVALID_REQUEST, "Invalid request body provided.")
		return
	}
	err = validateDeathNotice(deathNotice)
	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusBadRequest, INVALID_REQUEST, err.Error())
		return
	}

	err = validateDeathNotice(deathNotice)

	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusBadRequest, INVALID_REQUEST, err.Error())
		return
	}

	json.NewEncoder(w).Encode(deathNotice)
}

func validateDeathNotice(deathNotice createDeathNoticeRequest) error {
	errorMessages := make([]string, 5)

	if strings.TrimSpace(deathNotice.CauseOfDeath) == "" {
		errorMessages = append(errorMessages, "Cause of death cannot be empty")
	}

	if strings.TrimSpace(deathNotice.FirstName) == "" {
		errorMessages = append(errorMessages, "First name can not be empty")
	}

	if len(strings.TrimSpace(deathNotice.FirstName)) < 3 {
		errorMessages = append(errorMessages, "First name must be at least 3 characters long")
	}

	if strings.TrimSpace(deathNotice.LastName) == "" {
		errorMessages = append(errorMessages, "Last name can not be empty")
	}

	if len(strings.TrimSpace(deathNotice.LastName)) < 3 {
		errorMessages = append(errorMessages, "Last name must be at least 3 characters long")
	}

	if deathNotice.DateOfDeath.IsZero() {
		errorMessages = append(errorMessages, "Date of death cannot be empty")
	}

	if deathNotice.DateOfBirth.IsZero() {
		errorMessages = append(errorMessages, "Date of birth cannot be empty")
	}

	if deathNotice.DateOfBirth.After(deathNotice.DateOfDeath) {
		errorMessages = append(errorMessages, "Date of birth cannot be after date of death")
	}

	var filteredErrorMessages []string
	for _, msg := range errorMessages {
		if msg != "" {
			filteredErrorMessages = append(filteredErrorMessages, msg)
		}
	}

	if len(filteredErrorMessages) > 0 {
		return fmt.Errorf(strings.Join(filteredErrorMessages, ", "))
	}
	return nil
}

func (app *Application) GetDeathNoticeById(w http.ResponseWriter, r *http.Request) {
	var deathNotice createDeathNoticeRequest
	json.NewEncoder(w).Encode(deathNotice)
}

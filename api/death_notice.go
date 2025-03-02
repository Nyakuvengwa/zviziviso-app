package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type createNewFuneralParlourRequest struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	ContactNumber string `json:"contact_number"`
	Email         string `json:"email"`
}

type createNewAddressRequest struct {
	AddressType   string `json:"address_type"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Province      string `json:"province"`
	PostalCode    string `json:"postal_code"`
	ContactPerson string `json:"contact_person"`
	ContactNumber string `json:"contact_number"`
}

type createDeathNoticeRequest struct {
	FullName       string                         `json:"full_name"`
	DateOfDeath    time.Time                      `json:"date_of_death"`
	Age            int32                          `json:"age"`
	CauseOfDeath   string                         `json:"cause_of_death"`
	FuneralParlour createNewFuneralParlourRequest `json:"funeral_parlour"`
	Address        createNewAddressRequest        `json:"address"`
	Obituary       string                         `json:"obituary"`
	ImageUrl       string                         `json:"image_url"`
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

	// TODO: Add logic which saves the Funeral parlour, Address, notice in that order and in a transaction.

	json.NewEncoder(w).Encode(deathNotice)
}

func validateDeathNotice(deathNotice createDeathNoticeRequest) error {
	errorMessages := make([]string, 5)
	if deathNotice.FuneralParlour == (createNewFuneralParlourRequest{}) {

	}

	if len(errorMessages) > 0 {
		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}
	return nil
}

func (app *Application) NewDeathNotice(w http.ResponseWriter, r *http.Request) {
	var deathNotice createDeathNoticeRequest
	json.NewEncoder(w).Encode(deathNotice)
}

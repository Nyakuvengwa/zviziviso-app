package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (app *Application) ListCountries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	countries, err := app.service.ListCountries(ctx)
	if err != nil {
		log.Printf("Error: %v, ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(countries)
}

func (h *Application) GetCountryById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	countryIdString := r.PathValue("countryId")
	countryIdNormal, err := strconv.Atoi(countryIdString)
	if err != nil {
		NewProblemDetails(w, http.StatusBadRequest, "Invalid country id provided", "Invalid url path parameter provided.")
		return
	}
	countryId := int32(countryIdNormal)
	countries, err := h.service.GetCountry(ctx, countryId)
	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetails(w, http.StatusInternalServerError, "Unhandled server error", err.Error())
		return
	}
	json.NewEncoder(w).Encode(countries)
}

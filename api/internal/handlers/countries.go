package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	services "zviziviso-app/api/internal/services"
)

type CountryHandler struct {
	service services.CountryService
}

func NewCountryHandler(service services.CountryService) *CountryHandler {
	return &CountryHandler{service: service}
}

func (h *CountryHandler) ListCountries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	countries, err := h.service.ListCountries(ctx)
	if err != nil {
		log.Printf("Error: %v, ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(countries)
}

func (h *CountryHandler) GetCountryById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	countryIdString := r.PathValue("countryId")
	countryIdNormal, err := strconv.Atoi(countryIdString)
	if err != nil {
		pd := NewProblemDetails(http.StatusBadRequest, "Invalid country id provided", "Invalid url path parameter provided.")
		WriteProblemDetails(w, pd)
		return
	}
	countryId := int32(countryIdNormal)
	countries, err := h.service.GetCountryById(ctx, countryId)
	if err != nil {
		log.Printf("Error: %v, ", err)
		pd := NewProblemDetails(http.StatusInternalServerError, "Unhandled server error", err.Error())
		WriteProblemDetails(w, pd)
		return
	}
	json.NewEncoder(w).Encode(countries)
}

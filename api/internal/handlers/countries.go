package handlers

import (
	"encoding/json"
	"log"
	"net/http"
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
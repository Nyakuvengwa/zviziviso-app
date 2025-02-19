package routing

import (
	"net/http"
	"zviziviso-app/api/internal/handlers"
)

type CountryRoutes struct {
	handler *handlers.CountryHandler
}

func NewCountryRoutes(handler *handlers.CountryHandler) *CountryRoutes {
	return &CountryRoutes{handler: handler}
}

func (cr *CountryRoutes) SetupCountriesRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /countries", cr.handler.ListCountries)
}

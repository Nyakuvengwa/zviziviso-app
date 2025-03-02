package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zviziviso-app/internal/models"
)

var (
	INVALID_REQUEST       string = "Invalid request"
	INTERNAL_SERVER_ERROR string = "Internal server error"
)

func NewProblemDetailsErrorResponse(w http.ResponseWriter, status int, title string, detail string) {
	pd := models.ProblemDetails{
		Type:   fmt.Sprintf("%s/%d", "https://httpstatuses.com", status),
		Title:  title,
		Status: status,
		Detail: detail,
	}

	w.Header().Set("Content-Type", "application/problem+json") // Important!
	w.WriteHeader(pd.Status)
	json.NewEncoder(w).Encode(pd)
}

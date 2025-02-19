package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zviziviso-app/api/internal/models"
)

func WriteProblemDetails(w http.ResponseWriter, pd models.ProblemDetails) {
	w.Header().Set("Content-Type", "application/problem+json") // Important!
	w.WriteHeader(pd.Status)
	json.NewEncoder(w).Encode(pd)
}

func NewProblemDetails(status int, title string, detail string) models.ProblemDetails {
	return models.ProblemDetails{
		Type:   fmt.Sprintf("%s/%d", "https://httpstatuses.com", status),
		Title:  title,
		Status: status,
		Detail: detail,
	}
}

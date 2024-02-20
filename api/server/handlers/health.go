package handlers

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	data := ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Health Check",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

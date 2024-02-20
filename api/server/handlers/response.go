package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorDataResponse struct {
	Message string `json:"message"`
}

type successDataResponse struct {
	StatusCode string      `json:"statusCode"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
}

func errorResponse(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)

	response := errorDataResponse{
		Message: msg,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func successResponse(w http.ResponseWriter, op string, data interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	response := successDataResponse{
		StatusCode: fmt.Sprintf("%d", http.StatusOK),
		Message:    op,
		Data:       data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

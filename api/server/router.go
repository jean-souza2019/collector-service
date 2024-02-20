package server

import (
	"net/http"

	"github.com/jean-souza2019/collector-service/api/server/handlers"
)

func InitializeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/user", handlers.UserHandler)
	mux.HandleFunc("/user-find", handlers.FindUsersHandler)
}

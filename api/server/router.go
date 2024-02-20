package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jean-souza2019/collector-service/api/server/handlers"
)

func InitializeRoutes(r *mux.Router) {
	r.HandleFunc("/health", handlers.HealthHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", handlers.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/all", handlers.FindUsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/billing", handlers.CreateBillingHandler).Methods(http.MethodPost)
	r.HandleFunc("/billing/all", handlers.FindBillingsHandler).Methods(http.MethodGet)
}

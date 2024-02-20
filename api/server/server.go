package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Initialize() {
	r := mux.NewRouter()

	InitializeRoutes(r)

	fmt.Println("Server is running on :3000")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

package server

import (
	"fmt"
	"log"
	"net/http"
)

func Initialize() {
	mux := http.NewServeMux()

	InitializeRoutes(mux)

	fmt.Println("Server is running on :3000")

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

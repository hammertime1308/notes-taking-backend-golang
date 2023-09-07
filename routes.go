package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {

	// api/v1/ping route -> health check
	r.HandleFunc("/api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"status": "running"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

}

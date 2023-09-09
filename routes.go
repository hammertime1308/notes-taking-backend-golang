package main

import (
	"net/http"
	"notes-taking-backend-golang/pkg/controllers"

	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {

	// api/v1/ping route -> health check
	r.HandleFunc("/api/v1/ping", controllers.PingController).Methods(http.MethodGet)

}

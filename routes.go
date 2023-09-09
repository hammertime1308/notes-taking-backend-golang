package main

import (
	"net/http"
	"notes-taking-backend-golang/pkg/controllers"
	"notes-taking-backend-golang/pkg/repository"

	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router, db repository.Repository) {
	controller := controllers.NewController(db)

	// api/v1/ping route -> health check
	r.HandleFunc("/api/v1/ping", controller.PingController).Methods(http.MethodGet)

	// api/v1/signup route -> register a new user
	r.HandleFunc("/api/v1/signup", controller.SignUp).Methods(http.MethodPost)
}

package controllers

import (
	"net/http"
	"notes-taking-backend-golang/pkg/repository"

	"github.com/gorilla/mux"
)

type controller struct {
	repository repository.Repository
}

func NewController(repo repository.Repository) *controller {
	return &controller{
		repository: repo,
	}
}

func (c *controller) RegisterRoutes(r *mux.Router) {

	// api/v1/ping route -> health check
	r.HandleFunc("/api/v1/ping", c.pingController).Methods(http.MethodGet)

	// api/v1/signup route -> register a new user
	r.HandleFunc("/api/v1/signup", c.signUp).Methods(http.MethodPost)

	// api/v1/login -> login the user and returns
	r.HandleFunc("/api/v1/login", c.login).Methods(http.MethodPost)

	// api/v1/notes -> adds new note against the user
	r.HandleFunc("/api/v1/notes", c.addNote).Methods(http.MethodPost)
}

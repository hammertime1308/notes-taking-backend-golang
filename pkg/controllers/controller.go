package controllers

import "notes-taking-backend-golang/pkg/repository"

type controller struct {
	repository repository.Repository
}

func NewController(repo repository.Repository) *controller {
	return &controller{
		repository: repo,
	}
}

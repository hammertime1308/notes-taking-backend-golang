package repository

import (
	"context"
	"notes-taking-backend-golang/models"
)

type Repository interface {
	Connect() error
	Close() error
	AddNewUser(ctx context.Context, user models.User) error
	Login(ctx context.Context, user models.User) (string, error)
	AddNote(ctx context.Context, note models.Note) (models.Note, error)
	GetAllNotes(ctx context.Context, user models.User) ([]models.Note, error)
	DeleteNote(ctx context.Context, note models.Note) error
}

func New(config *models.Config) Repository {
	return &sql{
		username: config.Database.Username,
		password: config.Database.Password,
		hostname: config.Database.Host,
		port:     config.Database.Port,
		dbName:   config.Database.DbName,
	}
}

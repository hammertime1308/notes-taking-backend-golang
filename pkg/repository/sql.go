package repository

import (
	"context"
	"fmt"
	"notes-taking-backend-golang/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sql struct {
	*sqlx.DB
	username string
	password string
	hostname string
	port     int
	dbName   string
}

const (
	ADD_USER = `INSERT INTO users(name,email,password,session_id) VALUES(?,?,?,?)`
)

func (s *sql) Connect() error {
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", s.username, s.password, s.hostname, s.port, s.dbName)

	var err error
	s.DB, err = sqlx.Connect("mysql", dataSourceName)
	return err
}

func (s *sql) Close() error {
	return s.DB.Close()
}

func (s *sql) AddNewUser(ctx context.Context, user models.User) error {
	_, err := s.ExecContext(ctx, ADD_USER, user.Name, user.Email, user.Password, user.SessionId)
	return err
}

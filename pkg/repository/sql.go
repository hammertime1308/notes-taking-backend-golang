package repository

import (
	"context"
	"errors"
	"fmt"
	"notes-taking-backend-golang/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
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
	ADD_USER          = `INSERT INTO users(name,email,password) VALUES(?,?,?)`
	CHECK_USER_EXISTS = `SELECT user_id FROM users WHERE email = ? AND password = ? LIMIT 1`
	ADD_USER_SESSION  = `INSERT INTO user_session(user_id,session_id) VALUES (?,?) ON DUPLICATE KEY UPDATE session_id=VALUES(session_id)`
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
	_, err := s.ExecContext(ctx, ADD_USER, user.Name, user.Email, user.Password)
	return err
}

func (s *sql) Login(ctx context.Context, user models.User) (string, error) {
	exists, userId, err := s.checkUserExists(ctx, user)
	if err != nil {
		return user.SessionId, err
	}

	if !exists {
		return user.SessionId, errors.New("user does not exist")
	}

	_, err = s.ExecContext(ctx, ADD_USER_SESSION, userId, user.SessionId)
	return user.SessionId, err
}

func (s *sql) checkUserExists(ctx context.Context, user models.User) (bool, int64, error) {
	var userId null.Int

	err := s.GetContext(ctx, &userId, CHECK_USER_EXISTS, user.Email, user.Password)
	return userId.Valid, userId.Int64, err
}

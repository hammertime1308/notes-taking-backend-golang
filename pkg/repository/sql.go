package repository

import (
	"fmt"

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

func (s *sql) Connect() error {
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", s.username, s.password, s.hostname, s.port, s.dbName)

	var err error
	s.DB, err = sqlx.Connect("mysql", dataSourceName)
	return err
}

func (s *sql) Close() error {
	return s.Close()
}

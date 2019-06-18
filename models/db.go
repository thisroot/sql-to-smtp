package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type Datastore interface {
	AllMails() ([]*Mail, error)
}

type DB struct {
	*sql.DB
}

func NewDB(config *mysql.Config) (*DB, error) {
	logrus.Info("connect to: ", config.FormatDSN())
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
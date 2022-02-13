package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	conn *sql.DB
}

func NewDataBase(path string) (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	} else {
		return &Database{db}, nil
	}
}

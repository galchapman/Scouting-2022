package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func NewDataBase(path string) (*Database, error) {
	sqlDB, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db := Database{sqlDB}

	err = db.init()

	if err != nil {
		return nil, err
	}

	return &db, nil
}

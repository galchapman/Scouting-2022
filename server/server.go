package server

import "Scouting-2022/server/database"

type Server struct {
	db *database.Database
}

func NewServer() (*Server, error) {
	db, err := database.NewDataBase("./database.sqlite")
	if err != nil {
		return nil, err
	} else {
		return &Server{db}, nil
	}
}

package server

import (
	"Scouting-2022/server/database"
	"Scouting-2022/server/toa_api"
)

type Server struct {
	db     *database.Database
	client *toa_api.TOAClient
	users  map[int]database.User
	event  Event
}

func NewServer(TOAApiKey string, eventKey string) (*Server, error) {
	var self *Server
	var db *database.Database
	var err error
	db, err = database.NewDataBase("./database.sqlite3")
	if err != nil {
		return nil, err
	}
	self = &Server{db: db, client: toa_api.NewTOAClient(TOAApiKey, "Megiddo Lions Scouting System"),
		users: make(map[int]database.User)}
	// get event info from The orange alliance api
	self.event, err = self.getEvent(eventKey)
	if err != nil {
		return nil, err
	}

	return self, nil
}

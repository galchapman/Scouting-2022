package server

import (
	"Scouting-2022/server/database"
	"Scouting-2022/server/toa_api"
	"net/http"
)

type Server struct {
	db      *database.Database
	client  *toa_api.TOAClient
	users   map[int]database.User
	event   Event
	http    http.Server
	servMux *http.ServeMux
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
		users:   make(map[int]database.User),
		http:    http.Server{Addr: ":80"},
		servMux: http.NewServeMux()}
	// get event info from The orange alliance api
	self.event, err = self.getEvent(eventKey)
	if err != nil {
		return nil, err
	}

	// config
	self.configHTTP()

	return self, nil
}

func (server *Server) configHTTP() {
	server.http.Handler = server.servMux

	server.servMux.HandleFunc("/robots", server.handleRobots)
}

func (server *Server) Run() error {
	println("Server running at localhost:80")
	return server.http.ListenAndServe()
}

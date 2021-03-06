package server

import (
	"Scouting-2022/server/database"
	"Scouting-2022/server/toa_api"
	"net/http"
	"sync"

	"github.com/NYTimes/gziphandler"
)

type Server struct {
	db       *database.Database
	client   *toa_api.TOAClient
	sessions map[string]Session
	http     http.Server
	servMux  *http.ServeMux
	mu       sync.Mutex
}

func NewServer(TOAApiKey string, eventKey string) (*Server, error) {
	var self *Server
	var db *database.Database
	var err error
	db, err = database.NewDataBase("./database.sqlite3")
	if err != nil {
		return nil, err
	}
	_ = db.NewUser("Admin", "password", "Admin", "Admin")

	self = &Server{db: db, // client: toa_api.NewTOAClient(TOAApiKey, "Megiddo Lions Scouting System"),
		sessions: make(map[string]Session),
		http:     http.Server{Addr: ":80"},
		servMux:  http.NewServeMux()}
	// get event info from The orange alliance api
	// _, err = self.getEvent(eventKey)
	if err != nil {
		return nil, err
	}

	// config
	self.configHTTP()

	return self, nil
}

func serveFile(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "www"+req.URL.Path)
}

func (server *Server) configHTTP() {
	server.http.Handler = server.servMux

	server.servMux.HandleFunc("/", server.handleIndex)
	server.servMux.HandleFunc("/index.html", server.handleIndex)
	server.servMux.HandleFunc("/robots.html", server.handleRobots)
	server.servMux.HandleFunc("/main.css", serveFile)
	server.servMux.HandleFunc("/favicon.ico", serveFile)
	server.servMux.HandleFunc("/login.html", server.handleLogin)
	server.servMux.HandleFunc("/users.html", server.handleUsers)
	server.servMux.HandleFunc("/management.html", server.handleManagement)
	server.servMux.HandleFunc("/form.html", server.handleForm)
	server.servMux.HandleFunc("/post-form.html", server.handlePostForm)
	server.servMux.HandleFunc("/assign.html", server.handleAssign)
	server.servMux.HandleFunc("/create-user.html", server.handleCreateUser)
	server.servMux.HandleFunc("/match-data.html", server.handleMatchData)
	server.servMux.HandleFunc("/team.html", server.handleTeamPage)
	server.servMux.HandleFunc("/ranking.html", server.handleRanking)
	server.servMux.HandleFunc("/user-management.html", server.handleUserManagement)
	server.servMux.HandleFunc("/pit-form.html", server.handlePitForm)
	server.servMux.HandleFunc("/schedule.html", server.handleSchedule)
	server.servMux.HandleFunc("/sql.html", server.handleRunSQL)

	// API
	server.servMux.HandleFunc("/api/get_team_games", server.handleGetTeamGames)

	server.http.Handler = gziphandler.GzipHandler(server.http.Handler)
}

func (server *Server) Run() error {
	println("Server running at localhost:80")
	return server.http.ListenAndServe()
}

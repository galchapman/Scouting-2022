package server

import (
	"Scouting-2022/server/database"
	"net/http"
	"strconv"
)

var currentGame int
var currentGameScouters int

func (server *Server) handleManagement(w http.ResponseWriter, req *http.Request) {
	// check role
	if _, session := server.checkSession(req); session == nil || session.user.Role < database.ManagerRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	query := req.URL.Query()
	if query.Has("start") {
		game, err := strconv.Atoi(query.Get("start"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		currentGame = game
		currentGameScouters = 0
	} else {
		http.ServeFile(w, req, "www/management.html")
	}
}

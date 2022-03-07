package server

import (
	"Scouting-2022/server/database"
	"encoding/json"
	"net/http"
)

func (server *Server) handleGetTeamGames(w http.ResponseWriter, req *http.Request) {

	if _, session := server.checkSession(req); session == nil || session.user.Role < database.ViewerRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	var data []byte
	answers, err := server.db.GetTeamsGames()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err = json.Marshal(answers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

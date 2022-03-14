package server

import (
	"Scouting-2022/server/database"
	"net/http"
	"strconv"
)

func (server *Server) handlePitForm(w http.ResponseWriter, req *http.Request) {
	if _, session := server.checkSession(req); session == nil || session.user.Role < database.SupervisorRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	switch req.Method {
	case http.MethodGet:
		http.ServeFile(w, req, "www/pit-form.html")
		break
	case http.MethodPost:
		err := req.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !req.PostForm.Has("team") || !req.PostForm.Has("notes") {
			http.Error(w, "Invalid form format", http.StatusBadRequest)
			return
		}

		var team database.Team
		if team.TeamNumber, err = strconv.Atoi(req.PostForm.Get("team")); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Check team exits
		if _, err = server.db.GetTeam(team.TeamNumber); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// Insert notes
		err = server.db.SetTeamNotes(team, req.PostForm.Get("notes"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Serve file again
		http.ServeFile(w, req, "www/pit-form.html")
		break
	}
}

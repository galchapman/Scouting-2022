package server

import (
	"Scouting-2022/server/database"
	"html"
	"net/http"
	"os"
	"strconv"
)

var pitFormHtml string

func (server *Server) handlePitForm(w http.ResponseWriter, req *http.Request) {
	if _, session := server.checkSession(req); session == nil || session.user.Role < database.SupervisorRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	switch req.Method {
	case http.MethodGet:
		if err := req.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.Form.Has("team") {
			if pitFormHtml == "" {
				content, err := os.ReadFile("www/pit-form.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				pitFormHtml = string(content)
			}
			var err error
			var team database.Team
			var notes string
			if team.TeamNumber, err = strconv.Atoi(req.Form.Get("team")); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if team, err = server.db.GetTeam(team.TeamNumber); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if notes, err = server.db.GetTeamNotes(team); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var values map[string]string
			values = make(map[string]string)

			values["${TEAM}"] = strconv.Itoa(team.TeamNumber)
			values["${NOTES}"] = html.EscapeString(notes)

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			text, err := replaceAll(pitFormHtml, values)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, _ = w.Write([]byte(text))
			return
		}
		http.ServeFile(w, req, "www/pit-form.html")
		break
	case http.MethodPost:
		var err error
		if err = req.ParseForm(); err != nil {
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

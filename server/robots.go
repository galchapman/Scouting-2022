package server

import (
	"Scouting-2022/server/database"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (server *Server) handleRobots(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		query := req.URL.Query()
		if query.Has("team") {
			http.ServeFile(w, req, "robots/"+query.Get("team")+".jpeg")
			return
		} else if query.Has("missing") {
			teams, _ := server.db.GetTeams()
			var data string
			for _, team := range teams {
				if _, err := os.Stat("robots/" + strconv.Itoa(team.TeamNumber) + ".jpeg"); errors.Is(err, os.ErrNotExist) {
					data += team.Name + " - " + strconv.Itoa(team.TeamNumber) + "<br>"
				}
			}
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(data))
			return
		} else {
			if _, session := server.checkSession(req); session == nil || session.user.Role < database.SupervisorRole {
				http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
				return
			}
			http.ServeFile(w, req, "www/robots.html")
			return
		}
	case http.MethodPost:
		if _, session := server.checkSession(req); session == nil || session.user.Role < database.SupervisorRole {
			http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
			return
		}
		err := req.ParseMultipartForm(1 << 25)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file, _, err := req.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, _ := io.ReadAll(file)

		var filename []string
		var ok bool
		if filename, ok = req.MultipartForm.Value["group-number"]; !ok {
			http.Error(w, "No team number provided", http.StatusBadRequest)
			return
		}

		if strings.Count(filename[0], "..") != 0 {
			http.Error(w, "File name cannot contain ..", http.StatusBadRequest)
			return
		}

		_ = os.WriteFile("robots/"+filename[0]+".jpeg", data, os.ModePerm)

		w.Header().Set("Location", "robots.html")
		w.WriteHeader(http.StatusSeeOther)
		break
	default:
		http.Error(w, "You know what you did", 405)
	}
}

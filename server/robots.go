package server

import (
	"Scouting-2022/server/database"
	"io"
	"net/http"
	"os"
	"strings"
)

func (server *Server) handleRobots(w http.ResponseWriter, req *http.Request) {

	if _, session := server.checkSession(req); session == nil || session.user.Role < database.SupervisorRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	switch req.Method {
	case http.MethodGet:
		query := req.URL.Query()
		if query.Has("team") {
			http.ServeFile(w, req, "robots/"+query.Get("team")+".jpeg")
		} else {
			http.ServeFile(w, req, "www/robots.html")
		}
		break
	case http.MethodPost:
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

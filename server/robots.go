package server

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func (server *Server) handleRobots(w http.ResponseWriter, req *http.Request) {
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
		file, header, err := req.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, _ := io.ReadAll(file)

		if strings.Count(header.Filename, "..") != 0 {
			http.Error(w, "File name cannot contain ..", http.StatusBadRequest)
			return
		}

		_ = os.WriteFile("robots/"+header.Filename, data, os.ModePerm)

		w.Header().Set("Location", "robots.html")
		w.WriteHeader(http.StatusSeeOther)
		break
	default:
		http.Error(w, "You know what you did", 405)
	}
}

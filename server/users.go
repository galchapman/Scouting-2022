package server

import (
	"Scouting-2022/server/database"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"
)

var usersHtml string

func (server *Server) handleUsers(w http.ResponseWriter, req *http.Request) {

	if _, session := server.checkSession(req); session == nil || session.user.Role < database.ManagerRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	if usersHtml == "" {
		content, err := os.ReadFile("www/users.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		usersHtml = string(content)
	}

	scouters, err := server.db.GetScouters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userTable string

	for index, scouter := range scouters {
		userTable += fmt.Sprintf("<tr onclick=\"SelectUser(%d)\"><td>%d</td><td>%s</td>", index, scouter.ID, html.EscapeString(scouter.Name))
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(strings.Replace(usersHtml, "${USERS}", userTable, 1)))
}

func (server *Server) handleCreateUser(w http.ResponseWriter, req *http.Request) {
	if _, session := server.checkSession(req); session == nil || session.user.Role < database.ManagerRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	switch req.Method {
	case http.MethodGet:
		http.ServeFile(w, req, "www/create-user.html")
		break
	case http.MethodPost:
		err := req.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = server.db.NewUser(req.PostForm.Get("username"), req.PostForm.Get("password"), req.PostForm.Get("screen-name"), req.PostForm.Get("role"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			http.ServeFile(w, req, "www/create-user.html")
		}
		break
	}
}

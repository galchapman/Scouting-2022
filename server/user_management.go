package server

import (
	"Scouting-2022/server/database"
	"fmt"
	"html"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var userManagementHtml string

func (server *Server) handleUserManagement(w http.ResponseWriter, req *http.Request) {
	if userManagementHtml == "" {
		content, err := os.ReadFile("www/user-management.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		userManagementHtml = string(content)
	}

	// check role
	if _, session := server.checkSession(req); session == nil || session.user.Role < database.ManagerRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	switch req.Method {
	case http.MethodGet:
		var usersTable string
		var users []database.User
		var err error
		users, err = server.db.GetUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, user := range users {
			usersTable += fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%s</td><td>%d</td></tr>", user.ID, html.EscapeString(user.Name), html.EscapeString(user.ScreenName), user.Role)
		}

		w.Header().Set("Content-type", "text/html; encoding: utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(strings.Replace(userManagementHtml, "${USERS}", usersTable, 1)))
		break
	case http.MethodPost:
		err := req.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !req.PostForm.Has("action") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		switch req.PostForm.Get("action") {
		case "delete":
			var user database.User
			user.ID, err = strconv.Atoi(req.PostForm.Get("user"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if !req.PostForm.Has("action") {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = server.db.DeleteUser(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(""))
			return
		case "modify":
			var user database.User
			user.ID, err = strconv.Atoi(req.PostForm.Get("user"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if !req.PostForm.Has("action") {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if !req.PostForm.Has("name") {
				http.Error(w, "name has to be", http.StatusBadRequest)
				return
			}
			if !req.PostForm.Has("role") {
				http.Error(w, "role has to be", http.StatusBadRequest)
				return
			}

			user.ScreenName = req.PostForm.Get("name")
			user.Role = database.ParseRole(req.PostForm.Get("role"))
			if user.Role == -1 {
				http.Error(w, "Invalid role", http.StatusBadRequest)
				return
			}

			err = server.db.ModifyUser(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(""))
			return
		}
	}
}

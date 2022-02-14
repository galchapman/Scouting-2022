package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var usersHtml string

func (server *Server) handleUsers(w http.ResponseWriter, req *http.Request) {
	if usersHtml == "" {
		content, err := os.ReadFile("www/users.html")
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			println("ERROR index:15 " + err.Error())
			return
		}
		usersHtml = string(content)
	}

	scouters, err := server.db.GetScouters()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		println("ERROR index:15 " + err.Error())
		return
	}

	var userTable string

	for index, scouter := range scouters {
		userTable += fmt.Sprintf("<tr onclick=\"SelectUser(%d)\"><td>%d</td><td>%s</td>", index, scouter.ID, scouter.Name)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(strings.Replace(usersHtml, "${USERS}", userTable, 1)))
}

package server

import (
	"Scouting-2022/server/database"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var scheduleHtml string

func (server *Server) handleSchedule(w http.ResponseWriter, req *http.Request) {
	var err error
	if err = req.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if scheduleHtml == "" {
		content, err := os.ReadFile("www/schedule.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		scheduleHtml = string(content)
	}

	if !req.Form.Has("team") {
		http.Error(w, "must include a team", http.StatusBadRequest)
		return
	}

	var team database.Team
	if team.TeamNumber, err = strconv.Atoi(req.Form.Get("team")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var schedule []database.Game
	if schedule, err = server.db.GetTeamSchedule(currentGame, team); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data string

	for _, game := range schedule {
		data += fmt.Sprintf("<tr><td>%d</td><td>%d</td><td>%d</td><td>%d</td><td>%d</td></tr>", game.ID, game.Red1.TeamNumber, game.Red2.TeamNumber, game.Blue1.TeamNumber, game.Blue2.TeamNumber)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(strings.Replace(scheduleHtml, "${DATA}", data, 1)))
}

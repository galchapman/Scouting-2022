package server

import (
	"Scouting-2022/server/database"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var matchDataHtml string

func (server *Server) handleMatchData(w http.ResponseWriter, req *http.Request) {

	if _, session := server.checkSession(req); session == nil || session.user.Role < database.ViewerRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	if matchDataHtml == "" {
		content, err := os.ReadFile("www/match-data.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		matchDataHtml = string(content)
	}

	games, err := server.db.GetTeamsGames()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data string

	for _, game := range games {
		data += fmt.Sprintf("<tr><td>%d</td><td>%d</td><td>%d</td><td>%s</td><td>%s</td><td>%s</td><td>%d</td><td>%d</td><td>%d</td><td>%s</td><td>%d</td><td>%d</td><td>%d</td><td>%d</td><td>%d</td><td>%d</td><td>%s</td><td>%s</td><td>%s</td></tr>",
			game.Game.ID, game.Scouter.ID, game.Team.TeamNumber, game.Location, game.Alliance, game.AutoDuck, game.AutoStorage, game.AutoShipping, game.CubeLevel, game.Parking, game.Storage, game.ShippingLow, game.ShippingMid, game.ShippingHigh, game.Shared, game.TeleopDucks, game.Capping, game.Worked, game.Notes)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(strings.Replace(matchDataHtml, "${DATA}", data, 1)))
}

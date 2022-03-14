package server

import (
	"Scouting-2022/server/database"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var rankingHtml string

func (server *Server) handleRanking(w http.ResponseWriter, req *http.Request) {

	if _, session := server.checkSession(req); session == nil || session.user.Role < database.ViewerRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	if rankingHtml == "" {
		content, err := os.ReadFile("www/ranking.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		rankingHtml = string(content)
	}

	var err error
	var games []database.FormAnswer
	var scores map[int]database.TeamScore

	err = req.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	games, err = server.db.GetTeamsGames()
	scores = make(map[int]database.TeamScore)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Calculate scores
	for _, game := range games {
		score, _ := scores[game.Team.TeamNumber]
		score.Add(game)
		scores[game.Team.TeamNumber] = score
	}
	// Write them
	var data string

	for team, score := range scores {
		s := score
		data += fmt.Sprintf("<tr><td>%d</td><td><a href=\"/team.html?team=%d\">%d</a></td><td>%.2f</td><td>%.2f</td><td>%.2f</td><td>%.2f</td><td>%.2f</td><td>%.2f</td><td>%.2f</td><td>%.2f</td></tr>",
			s.GamesPlayed, team, team, float32(s.TotalScore)/float32(s.GamesPlayed), float32(s.AutoTotalScore)/float32(s.GamesPlayed),
			float32(s.TotalDucksCountAuto)/float32(s.AutoStartedNear), float32(s.TotalTowerScoreAuto)/float32(s.AutoStartedNear),
			float32(s.TotalTowerScore)/float32(s.GamesPlayed), float32(s.TotalSharedScore)/float32(s.GamesPlayed),
			float32(s.TotalDucksCount)/float32(s.GamesPlayed), float32(s.TotalShippingElementsPlaced)/float32(s.GamesPlayed))
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(strings.Replace(rankingHtml, "${DATA}", data, 1)))
}

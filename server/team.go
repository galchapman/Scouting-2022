package server

import (
	"Scouting-2022/server/database"
	"fmt"
	"html"
	"net/http"
	"os"
	"strconv"
)

var teamPageHtml string

func (server *Server) handleTeamPage(w http.ResponseWriter, req *http.Request) {

	if _, session := server.checkSession(req); session == nil || session.user.Role < database.ViewerRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	if teamPageHtml == "" {
		content, err := os.ReadFile("www/team-page.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		teamPageHtml = string(content)
	}

	var err error
	var values map[string]string
	var team database.Team
	var games []database.FormAnswer
	var superForms []database.TeamSupervisorForm
	var score database.TeamScore

	values = make(map[string]string)

	err = req.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	team.TeamNumber, err = strconv.Atoi(req.Form.Get("team"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	values["${GROUP_NUMBER}"] = strconv.Itoa(team.TeamNumber)

	team, err = server.db.GetTeam(team.TeamNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	values["${TEAM_NAME}"] = team.Name

	games, err = server.db.GetTeamGames(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	superForms, err = server.db.GetTeamSuperForms(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var notes string
	for _, game := range games {
		score.Add(game)
		notes += html.EscapeString(game.Notes) + "<br>"
	}

	values["${NOTES}"] = notes
	values["${AVERAGE_TOTAL}"] = strconv.Itoa(score.TotalDucksCount)
	values["${AVERAGE_WORK}"] = fmt.Sprintf("%.2f", float32(score.Worked)/float32(len(games)))
	values["${AVERAGE_AUTO}"] = fmt.Sprintf("%.2f", float32(score.AutoTotalScore)/float32(len(games)))
	values["${PERCENT_DUCK}"] = fmt.Sprintf("%.2f", 100*float32(score.TotalDucksCountAuto)/float32(score.AutoStartedNear))
	values["${AVERAGE_ELEMENTS_AUTO}"] = fmt.Sprintf("%.2f", float32(score.TotalTowerScoreAuto)/float32(score.AutoStartedNear))
	values["${AVERAGE_SHARED}"] = fmt.Sprintf("%.2f", float32(score.TotalSharedScore)/float32(score.GamesPlayed))
	values["${AVERAGE_TOWER}"] = fmt.Sprintf("%.2f", float32(score.TotalTowerScore)/float32(score.GamesPlayed))
	values["${AVERAGE_DUCKS}"] = fmt.Sprintf("%.2f", float32(score.TotalDucksCount)/float32(score.GamesPlayed))
	values["${PERCENT_CUP}"] = fmt.Sprintf("%.2f", 100*float32(score.TotalShippingElementsPlaced)/float32(score.GamesPlayed))
	values["${MATCHES}"] = strconv.Itoa(score.GamesPlayed)

	var data string

	for _, game := range games {
		data += fmt.Sprintf("<tr><td>%d</td><td>%d</td><td>%d</td><td>%s</td><td>%s</td><td>%s</td><td>%d</td><td>%d</td><td>%d</td><td>%s</td><td>%d</td><td>%d</td><td>%d</td><td>%d</td><td>%d</td><td>%d</td><td>%s</td><td>%s</td><td>%s</td></tr>",
			game.Game.ID, game.Scouter.ID, game.Team.TeamNumber, game.Location, game.Alliance, game.AutoDuck, game.AutoStorage, game.AutoShipping, game.CubeLevel, game.Parking, game.Storage, game.ShippingLow, game.ShippingMid, game.ShippingHigh, game.Shared, game.TeleopDucks, game.Capping, game.Worked, game.Notes)
	}
	values["${DATA}"] = data

	var superNotes string
	var fauls int
	for _, superForm := range superForms {
		superNotes += html.EscapeString(superForm.Notes) + "<br>"
		switch superForm.Penalty {
		case "Red":
			fauls += 2
			break
		case "Yellow":
			fauls++
			break
		}
	}

	values["${SUPER_NOTES}"] = superNotes
	values["${FAULS}"] = strconv.Itoa(fauls)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	text, err := replaceAll(teamPageHtml, values)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write([]byte(text))
}

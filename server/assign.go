package server

import (
	"Scouting-2022/server/database"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var assignHtml string

func (server *Server) handleAssign(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if assignHtml == "" {
			content, err := os.ReadFile("www/assign.html")
			if err != nil {
				http.Error(w, "Not Found", 404)
				println("ERROR assign:24 " + err.Error())
				return
			}
			assignHtml = string(content)
		}

		games, err := server.db.GetGames()
		if err != nil {
			http.Error(w, "Not Found", 404)
			println("ERROR assign:24 " + err.Error())
			return
		}

		var gamesTable string

		for _, game := range games {
			gamesTable += fmt.Sprintf("<tr><td>%d</td><td>%d</td><td>%s</td><td>%d</td><td>%s</td><td>%d</td><td>%s</td><td>%d</td><td>%s</td></tr>",
				game.ID, game.Red1.TeamNumber, game.ScouterRed1.Name, game.Red2.TeamNumber, game.ScouterRed2.Name, game.Blue1.TeamNumber, game.ScouterBlue1.Name, game.Blue2.TeamNumber, game.ScouterBlue2.Name)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(strings.Replace(assignHtml, "${DATA}", gamesTable, 1)))

		break
	case http.MethodPost:
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		for _, line := range strings.Split(string(body), "\n") {
			var game database.Game
			var Red1 int
			var Red2 int
			var Blue1 int
			var Blue2 int
			var ScoutRed1 string
			var ScoutRed2 string
			var ScoutBlue1 string
			var ScoutBlue2 string
			_, err = fmt.Sscanf(line, "%d,%d,%s,%d,%s,%d,%s,%d,%s,", &game.ID, &Red1, &ScoutRed1, &Red2, &ScoutRed2, &Blue1, &ScoutBlue1, &Blue2, &ScoutBlue2)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				panic(err)
			}

			game.Red1, _ = server.event.groups[Red1]
			game.Red2, _ = server.event.groups[Red2]
			game.Blue1, _ = server.event.groups[Blue1]
			game.Blue2, _ = server.event.groups[Blue2]

			game.ScouterRed1, err = server.db.GetUserByName(ScoutRed1)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			game.ScouterRed2, err = server.db.GetUserByName(ScoutRed2)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			game.ScouterBlue1, err = server.db.GetUserByName(ScoutBlue1)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			game.ScouterBlue2, err = server.db.GetUserByName(ScoutBlue2)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			err = server.db.InsertGame(game)
			if err != nil {
				err = server.db.UpdateGame(game)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			w.WriteHeader(http.StatusOK)
		}

		break
	default:
		http.Error(w, "Invalid http method", http.StatusMethodNotAllowed)
	}
}

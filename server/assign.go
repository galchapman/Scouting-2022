package server

import (
	"Scouting-2022/server/database"
	"bytes"
	"encoding/csv"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var assignHtml string

func (server *Server) handleAssign(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if req.Header.Get("Content-Type") == "text/csv" {
			var table string
			table += "Match,Cube,Red Team 1,Scouter,Red Team 2,Scouter,Blue Team 1,Scouter,Blue Team 2,Scouter,\n"

			games, err := server.db.GetGames()
			if err != nil {
				http.Error(w, "Not Found", 404)
				println("ERROR assign:24 " + err.Error())
				return
			}

			for _, game := range games {
				table += fmt.Sprintf("%d,%d,%d,%s,%d,%s,%d,%s,%d,%s,\n",
					game.ID, game.GameType,
					game.Red1.TeamNumber, game.ScouterRed1.Name,
					game.Red2.TeamNumber, game.ScouterRed2.Name,
					game.Blue1.TeamNumber, game.ScouterBlue1.Name,
					game.Blue2.TeamNumber, game.ScouterBlue2.Name)
			}

			w.Header().Set("Content-Type", "text/csv")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(table))
		} else {
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
				gamesTable += fmt.Sprintf("<tr><td>%d</td><td>%d</td><td>%d</td><td>%s</td><td>%d</td><td>%s</td><td>%d</td><td>%s</td><td>%d</td><td>%s</td></tr>",
					game.ID, game.GameType,
					game.Red1.TeamNumber, html.EscapeString(game.ScouterRed1.Name),
					game.Red2.TeamNumber, html.EscapeString(game.ScouterRed2.Name),
					game.Blue1.TeamNumber, html.EscapeString(game.ScouterBlue1.Name),
					game.Blue2.TeamNumber, html.EscapeString(game.ScouterBlue2.Name))
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(strings.Replace(assignHtml, "${DATA}", gamesTable, 1)))
		}
		break
	case http.MethodPost:
		err := req.ParseMultipartForm(1 << 25)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file, _, err := req.FormFile("content")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, _ := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		table, _ := csv.NewReader(bytes.NewReader(body)).ReadAll()

		for _, line := range table[1:] {
			var game database.Game
			var Red1 int
			var Red2 int
			var Blue1 int
			var Blue2 int

			game.ID, err = strconv.Atoi(line[0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				panic(err)
			}
			game.GameType = line[1]

			Red1, err = strconv.Atoi(line[2])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				panic(err)
			}
			Red2, err = strconv.Atoi(line[4])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				panic(err)
			}
			Blue1, err = strconv.Atoi(line[6])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				panic(err)
			}
			Blue2, err = strconv.Atoi(line[8])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				panic(err)
			}

			var ok bool

			if game.Red1, ok = server.event.groups[Red1]; !ok {
				http.Error(w, strconv.Itoa(Red1)+" one Team was not Found", http.StatusNotFound)
				return
			}

			if game.Red2, ok = server.event.groups[Red2]; !ok {
				http.Error(w, strconv.Itoa(Red2)+" Team was not Found", http.StatusNotFound)
				return
			}

			if game.Blue1, ok = server.event.groups[Blue1]; !ok {
				http.Error(w, strconv.Itoa(Blue1)+" one Team was not Found", http.StatusNotFound)
				return
			}

			if game.Blue2, ok = server.event.groups[Blue2]; !ok {
				http.Error(w, strconv.Itoa(Blue2)+" Team was not Found", http.StatusNotFound)
				return
			}

			game.ScouterRed1, err = server.db.GetUserByName(line[3])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			game.ScouterRed2, err = server.db.GetUserByName(line[5])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			game.ScouterBlue1, err = server.db.GetUserByName(line[7])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			game.ScouterBlue2, err = server.db.GetUserByName(line[9])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = server.db.InsertGame(game)
			if err != nil {
				err = server.db.UpdateGame(game)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
		w.Header().Set("Location", "assign.html")
		w.WriteHeader(http.StatusSeeOther)

		break
	default:
		http.Error(w, "Invalid http method", http.StatusMethodNotAllowed)
	}
}

package server

import (
	"Scouting-2022/server/database"
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
)

var currentGame int
var currentGameScouters map[int]database.Team
var formServMutex sync.Mutex
var formHtml string
var postFormHtml string
var preFormHtml string

func (server *Server) handleManagement(w http.ResponseWriter, req *http.Request) {
	if currentGameScouters == nil {
		currentGameScouters = make(map[int]database.Team)
	}

	// check role
	if _, session := server.checkSession(req); session == nil || session.user.Role < database.ManagerRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	query := req.URL.Query()
	if query.Has("start") {
		var game database.Game
		gameID, err := strconv.Atoi(query.Get("start"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		formServMutex.Lock()
		defer formServMutex.Unlock()
		currentGame = gameID
		game, err = server.db.GetGame(gameID)
		for k := range currentGameScouters {
			delete(currentGameScouters, k)
		}

		currentGameScouters[game.ScouterRed1.ID] = game.Red1
		currentGameScouters[game.ScouterRed2.ID] = game.Red2
		currentGameScouters[game.ScouterBlue1.ID] = game.Blue1
		currentGameScouters[game.ScouterBlue2.ID] = game.Blue2
	} else {
		http.ServeFile(w, req, "www/start-game.html")
	}
}

func (server *Server) handleForm(w http.ResponseWriter, req *http.Request) {
	var session *Session
	var team database.Team
	var ok bool

	if _, session = server.checkSession(req); session == nil {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	if team, ok = currentGameScouters[session.user.ID]; !ok {
		game, err := server.db.GetNextGame(session.user, currentGame)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("ברכותי סיימת להיום"))
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if preFormHtml == "" {
			content, err := os.ReadFile("www/pre-form.html")
			if err != nil {
				http.Error(w, "Not Found", http.StatusNotFound)
				println("ERROR index:15 " + err.Error())
				return
			}
			preFormHtml = string(content)
		}

		values := make(map[string]string)
		values["${NEXT_MATCH}"] = strconv.Itoa(game.ID)

		if game.ScouterRed1.ID == session.user.ID {
			values["${GROUP_NUMBER}"] = strconv.Itoa(game.Red1.TeamNumber)
			values["${GROUP_COLOR}"] = "אדומה"
		} else if game.ScouterRed2.ID == session.user.ID {
			values["${GROUP_NUMBER}"] = strconv.Itoa(game.Red2.TeamNumber)
			values["${GROUP_COLOR}"] = "אדומה"
		} else if game.ScouterBlue1.ID == session.user.ID {
			values["${GROUP_NUMBER}"] = strconv.Itoa(game.Blue1.TeamNumber)
			values["${GROUP_COLOR}"] = "כחולה"
		} else {
			values["${GROUP_NUMBER}"] = strconv.Itoa(game.Blue2.TeamNumber)
			values["${GROUP_COLOR}"] = "כחולה"
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		text, err := replaceAll(preFormHtml, values)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write([]byte(text))
		return
	}

	formServMutex.Lock()
	defer formServMutex.Unlock()

	switch req.Method {
	case http.MethodGet:
		if formHtml == "" {
			content, err := os.ReadFile("www/scout-form.html")
			if err != nil {
				http.Error(w, "Not Found", http.StatusNotFound)
				println("ERROR index:15 " + err.Error())
				return
			}
			formHtml = string(content)
		}

		game, _ := server.db.GetGame(currentGame)

		values := make(map[string]string)
		values["${GROUP_NUMBER}"] = strconv.Itoa(team.TeamNumber)
		values["${USER}"] = session.user.ScreenName
		values["${MATCH}"] = strconv.Itoa(currentGame)
		if game.Red1.TeamNumber == team.TeamNumber || game.Red2.TeamNumber == team.TeamNumber {
			values["${GROUP_COLOR}"] = "אדומה"
		} else {
			values["${GROUP_COLOR}"] = "כחולה"
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		text, err := replaceAll(formHtml, values)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write([]byte(text))
		break
	case http.MethodPost:
		err := req.ParseForm()
		if err != nil {
			println("ERROR session:80 " + err.Error())
			http.Error(w, "Some error", 500)
			return
		}

		ScouterID := session.user.ID
		Team, _ := currentGameScouters[ScouterID]
		Game := currentGame
		var formAnswer database.FormAnswerResponse

		formAnswer, err = parseScoutFormAnswer(req.PostForm)
		if err != nil {
			http.Error(w, "Error parsing body "+err.Error(), http.StatusBadRequest)
			return
		}

		err = server.db.InsertAnswer(ScouterID, Team.TeamNumber, Game, formAnswer)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", "/post-form.html")
		w.WriteHeader(http.StatusSeeOther)
		break
	default:
		http.Error(w, "You know what you did", http.StatusMethodNotAllowed)
	}
}

func parseScoutFormAnswer(form url.Values) (database.FormAnswerResponse, error) {
	var answer database.FormAnswerResponse
	var err error

	if value, ok := form["location"]; ok {
		answer.Location = value[0]
	} else {
		return answer, errors.New("field `location` not found in body")
	}

	if value, ok := form["team_element"]; ok {
		answer.TeamElement = value[0]
	} else {
		return answer, errors.New("field `team_element` not found in body")
	}

	if value, ok := form["auto_duck"]; ok {
		answer.AutoDuck = value[0]
	} else {
		return answer, errors.New("field `auto_duck` not found in body")
	}

	if value, ok := form["auto_storage"]; ok {
		answer.AutoStorage, err = strconv.Atoi(value[0])
		if err != nil {
			return answer, err
		}
	} else {
		return answer, errors.New("field `auto_storage` not found in body")
	}

	if value, ok := form["auto_shipping"]; ok {
		answer.AutoShipping, err = strconv.Atoi(value[0])
		if err != nil {
			return answer, err
		}
	} else {
		return answer, errors.New("field `auto_shipping` not found in body")
	}

	if value, ok := form["cube_level"]; ok {
		answer.CubeLevel, err = strconv.Atoi(value[0])
		if err != nil {
			return answer, err
		}
	} else {
		return answer, errors.New("field `cube_level` not found in body")
	}

	if value, ok := form["parking"]; ok {
		answer.Parking = value[0]
	} else {
		return answer, errors.New("field `parking` not found in body")
	}

	if value, ok := form["storage"]; ok {
		answer.Storage, err = strconv.Atoi(value[0])
		if err != nil {
			return answer, err
		}
	} else {
		return answer, errors.New("field `storage` not found in body")
	}

	if value, ok := form["shipping_low"]; ok {
		answer.ShippingLow, err = strconv.Atoi(value[0])
		if err != nil {
			return answer, err
		}
	} else {
		return answer, errors.New("field `shipping_low` not found in body")
	}

	if value, ok := form["shipping_mid"]; ok {
		answer.ShippingMid, err = strconv.Atoi(value[0])
		if err != nil {
			return answer, err
		}
	} else {
		return answer, errors.New("field `shipping_mid` not found in body")
	}

	if value, ok := form["shipping_high"]; ok {
		answer.ShippingHigh, err = strconv.Atoi(value[0])
		if err != nil {
			return answer, err
		}
	} else {
		return answer, errors.New("field `shipping_high` not found in body")
	}

	if value, ok := form["shared"]; ok {
		answer.Shared, err = strconv.Atoi(value[0])
		if err != nil {
			return answer, err
		}
	} else {
		return answer, errors.New("field `shared` not found in body")
	}

	if value, ok := form["teleop_ducks"]; ok {
		answer.TeleopDucks, err = strconv.Atoi(value[0])
		if err != nil {
			return answer, err
		}
	} else {
		return answer, errors.New("field `teleop_ducks` not found in body")
	}

	if value, ok := form["capping"]; ok {
		answer.Capping = value[0]
	} else {
		return answer, errors.New("field `capping` not found in body")
	}

	if value, ok := form["worked"]; ok {
		answer.Worked = value[0]
	} else {
		return answer, errors.New("field `worked` not found in body")
	}

	answer.Notes = form.Get("notes")

	return answer, nil
}

func (server *Server) handlePostForm(w http.ResponseWriter, req *http.Request) {
	if postFormHtml == "" {
		content, err := os.ReadFile("www/post-form.html")
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			println("ERROR index:15 " + err.Error())
			return
		}
		postFormHtml = string(content)
	}

	var session *Session
	var game database.Game
	var err error

	if _, session = server.checkSession(req); session == nil {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	game, err = server.db.GetNextGame(session.user, currentGame)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ברכותי סיימת להיום"))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	values := make(map[string]string)
	values["${NEXT_MATCH}"] = strconv.Itoa(game.ID)

	if game.ScouterRed1.ID == session.user.ID {
		values["${GROUP_NUMBER}"] = strconv.Itoa(game.Red1.TeamNumber)
		values["${GROUP_COLOR}"] = "אדומה"
	} else if game.ScouterRed2.ID == session.user.ID {
		values["${GROUP_NUMBER}"] = strconv.Itoa(game.Red2.TeamNumber)
		values["${GROUP_COLOR}"] = "אדומה"
	} else if game.ScouterBlue1.ID == session.user.ID {
		values["${GROUP_NUMBER}"] = strconv.Itoa(game.Blue1.TeamNumber)
		values["${GROUP_COLOR}"] = "כחולה"
	} else {
		values["${GROUP_NUMBER}"] = strconv.Itoa(game.Blue2.TeamNumber)
		values["${GROUP_COLOR}"] = "כחולה"
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	text, err := replaceAll(postFormHtml, values)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write([]byte(text))
}

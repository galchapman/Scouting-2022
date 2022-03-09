package server

import (
	"Scouting-2022/server/database"
	"crypto/rand"
	"math/big"
	"net/http"
	"strconv"
)

type Session struct {
	user database.User
}

func (server *Server) handleLogin(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		query := req.URL.Query()
		if query.Has("session") {
			w.Header().Set("Set-Cookie", "session="+query.Get("session"))
			w.Header().Set("Location", "/")
			w.WriteHeader(http.StatusSeeOther)
		} else if query.Has("mint") {
			if _, session := server.checkSession(req); session == nil || session.user.Role < database.ManagerRole {
				http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
				return
			}

			session := GenerateUniqueSessionValue()
			UID, err := strconv.Atoi(query.Get("mint"))
			if err != nil {
				http.Error(w, "Invalid request parameters for this action", 400)
				println("ERROR session:32 " + err.Error())
				return
			}
			// Get user
			var user database.User
			user, err = server.db.GetUser(UID)
			if err != nil {
				http.Error(w, "Invalid request parameters for this action", 400)
				println("ERROR session:32 " + err.Error())
				return
			}
			server.addSession(session, user)
			// Returns session
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "text/plain")
			_, err = w.Write([]byte(session))
			if err != nil {
				println("ERROR session:55 " + err.Error())
			}
		} else {
			http.ServeFile(w, req, "www/login.html")
		}
		break
	case http.MethodPost:
		err := req.ParseForm()
		if err != nil {
			println("ERROR session:80 " + err.Error())
			http.Error(w, "Some error", 500)
			return
		}

		var user database.User
		username := req.FormValue("username")
		password := req.FormValue("password")

		user, err = server.db.GetUserByName(username)
		if err != nil {
			println("ERROR session:90")
			http.Error(w, err.Error(), 500)
			return
		}

		err = user.TryLoggingIn(password)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		session := GenerateUniqueSessionValue()
		server.addSession(session, user)

		w.Header().Set("Set-Cookie", "session="+session)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusSeeOther)
		_, err = w.Write(make([]byte, 0))
		println(user.Name + " has logged on with session: " + session)
		break
	default:
		http.Error(w, "You know what you did", 405)
	}
}

const SessionLength = 16

func GenerateUniqueSessionValue() string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, SessionLength)
	for i := 0; i < SessionLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}

func (server *Server) addSession(session string, user database.User) {
	defer server.mu.Unlock()
	server.mu.Lock()
	server.sessions[session] = Session{user: user}
}

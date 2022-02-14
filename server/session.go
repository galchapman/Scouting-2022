package server

import (
	"Scouting-2022/server/database"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

type Session struct {
	user *database.User
}

var autoLogin = regexp.MustCompile("login.html?session=(?P<Session>[a-zA-Z0-9]{16})")
var mintSession = regexp.MustCompile("login.html?mint=(?P<UID>[0-9]+)")
var loginHtml []byte

func (server *Server) handleLogin(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if autoLogin.MatchString(req.URL.String()) {
			w.Header().Set("Set-Cookie", string(autoLogin.FindSubmatch([]byte(req.URL.String()))[0]))
			w.Header().Set("Location", "/")
			w.WriteHeader(http.StatusSeeOther)
		} else if mintSession.MatchString(req.URL.String()) {
			session := GenerateUniqueSessionValue()
			UID, err := strconv.Atoi(string(mintSession.FindSubmatch([]byte(req.URL.String()))[0]))
			if err != nil {
				http.Error(w, "Invalid request parameters for this action", 400)
				println("ERROR session:32 " + err.Error())
				return
			}
			// Get user
			var user *database.User
			var ok bool
			if user, ok = server.users[UID]; !ok { // Check if user is already loaded
				var newUser database.User
				newUser, err = server.db.GetUser(UID)
				if err != nil {
					http.Error(w, "User is not in our database", 400)
					println("ERROR session:42 " + err.Error())
					return
				}
				server.users[UID] = &newUser
				user = &newUser
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
			if loginHtml == nil {
				content, err := os.ReadFile("www/login.html")
				if err != nil {
					http.Error(w, "Not Found", 404)
					println("ERROR session:62 " + err.Error())
					return
				}
				loginHtml = content
			}

			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(200)
			_, err := w.Write(loginHtml)
			if err != nil {
				println("ERROR session:72 " + err.Error())
				return
			}
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

		var pUser *database.User
		var ok bool
		if pUser, ok = server.users[user.ID]; !ok {
			server.users[user.ID] = &user
			pUser = &user
		}

		session := GenerateUniqueSessionValue()
		server.addSession(session, pUser)

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
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, SessionLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (server *Server) addSession(session string, user *database.User) {
	defer server.mu.Unlock()
	server.mu.Lock()
	server.sessions[session] = Session{user: user}
}

package server

import (
	"net/http"
	"os"
	"strings"
)

var indexHtml string

func (server *Server) handleIndex(w http.ResponseWriter, req *http.Request) {
	if indexHtml == "" {
		content, err := os.ReadFile("www/index.html")
		if err != nil {
			http.Error(w, "Not Found", 404)
			println("ERROR index:14 " + err.Error())
			return
		}
		indexHtml = string(content)
	}
	var body string

	hasCookie, session := server.checkSession(req)
	if hasCookie {
		if session != nil {
			body = "היי " + session.user.ScreenName
		} else {
			body = "העוגייה שלך תקולה, פנה בהקדם לקבלת חדשה"
			w.Header().Set("Set-Cookie", "session=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT")
		}
	} else {
		body = "אתה חסר כל עוגייה"
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write([]byte(strings.Replace(indexHtml, "${body}", body, 1)))
}

package server

import (
	"net/http"
	"strings"
)

func (server *Server) handleRobots(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "robots/"+strings.Split(req.URL.String(), "=")[1]+".jpeg")
}

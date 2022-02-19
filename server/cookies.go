package server

import "net/http"

// checkSession -> hasCookie session
func (server *Server) checkSession(req *http.Request) (bool, *Session) {
	session, err := req.Cookie("session")
	if err != nil {
		return false, nil
	}

	if session, ok := server.sessions[session.Value]; ok {
		return true, &session
	} else {
		return true, nil
	}
}

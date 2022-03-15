package server

import (
	"Scouting-2022/server/database"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (server *Server) handleRunSQL(w http.ResponseWriter, req *http.Request) {
	if _, session := server.checkSession(req); session == nil || session.user.Role < database.ManagerRole {
		http.Error(w, "You must be logged in to preform this action", http.StatusForbidden)
		return
	}

	switch req.Method {
	case http.MethodGet:
		http.ServeFile(w, req, "www/sql.html")
		break
	case http.MethodPost:
		var sqlQuery string
		if bytes, err := io.ReadAll(req.Body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			sqlQuery = string(bytes)
		}

		var result *sql.Rows
		var err error

		if result, err = server.db.Query(sqlQuery); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var columns []string
		var data string
		if columns, err = result.Columns(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data = "<table><tr><th>" + strings.Join(columns, "</th><th>") + "</th></tr>"

		values := make([]interface{}, len(columns))
		scanArgs := make([]interface{}, len(values))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		for result.Next() {
			if err = result.Scan(scanArgs...); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data += "<tr>"
			var element string
			// Print data
			for _, value := range values {
				switch value.(type) {
				case nil:
					element = "NULL"
				case []byte:
					element = string(value.([]byte))
				default:
					element = fmt.Sprint(value)
				}

				data += "<td>" + element + "</td>"
			}

			data += "</tr>"
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(data + "</table>"))
	}
}

package server

import "Scouting-2022/server/database"

type Event struct {
	users  map[int]database.User
	groups map[int]database.Team
}

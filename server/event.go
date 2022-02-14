package server

import (
	"Scouting-2022/server/database"
)

type Event struct {
	groups map[int]database.Team
}

func (server *Server) getEvent(eventKey string) (Event, error) {
	participants, err := server.client.GetEventParticipants(eventKey)
	if err != nil {
		return Event{}, err
	}

	event := Event{make(map[int]database.Team)}

	for _, participant := range participants {
		event.groups[participant.TeamNumber] = database.Team{TeamNumber: participant.TeamNumber, Name: participant.Team.TeamNameShort}
		_, _ = server.db.InsertTeam(participant.TeamNumber, participant.Team.TeamNameShort)
	}

	return event, err
}

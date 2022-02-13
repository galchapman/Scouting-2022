package toa_api

import (
	"encoding/json"
	"time"
)

type EventInfo struct {
	EventParticipantKey   string     `json:"event_participant_key"`
	SeasonKey             string     `json:"season_key"`
	RegionKey             string     `json:"region_key"`
	LeagueKey             string     `json:"league_key"`
	EventCode             string     `json:"event_code"`
	FirstEventCode        string     `json:"first_event_code"`
	EventTypeKey          string     `json:"event_type_key"`
	EventRegionNumber     int        `json:"event_region_number"`
	DivisionKey           int        `json:"division_key"`
	DivisionName          string     `json:"division_name"`
	EventName             string     `json:"event_name"`
	StartDate             time.Time  `json:"start_data"`
	EndDate               time.Time  `json:"end_date"`
	WeekKey               time.Month `json:"week_key"`
	City                  string     `json:"city"`
	StateProvince         string     `json:"state_prov"`
	Country               string     `json:"country"`
	Venue                 string     `json:"venue"`
	Website               string     `json:"website"`
	TimeZone              string     `json:"time_zone"`
	IsPublic              string     `json:"is_public"`
	ActiveTournamentLevel string     `json:"active_tournament_level"`
	AllianceCount         int        `json:"alliance_count"`
	FieldCount            int        `json:"field_count"`
	AdvanceSpots          int        `json:"advance_spots"`
	AdvanceEvent          string     `json:"advance_event"`
	DataSource            int        `json:"data_source"`
}

type EventParticipant struct {
	EventParticipantKey string `json:"event_participant_key"`
	EventKey            string `json:"event_key"`
	TeamKey             string `json:"team_key"`
	TeamNumber          int    `json:"team_number"`
	IsActive            bool   `json:"is_active"`
	CardStatus          string `json:"card_status"`
	Team                Team   `json:"team"`
}

func (c *TOAClient) GetEventParticipants(eventKey string) ([]EventParticipant, error) {
	var err error
	var body []byte
	var participant []EventParticipant

	body, err = c.Get(`event/` + eventKey + "/teams")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &participant)
	if err != nil {
		return nil, err
	}
	return participant, nil
}

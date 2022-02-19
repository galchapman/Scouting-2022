package toa_api

type Team struct {
	TeamKey       string `json:"team_key"`
	RegionKey     string `json:"region_key"`
	LeagueKey     string `json:"league_key"`
	TeamNumber    int    `json:"team_number"`
	TeamNameShort string `json:"team_name_short"`
	TeamNameLong  string `json:"team_name_long"`
	RobotName     string `json:"robot_name"`
	LastActive    string `json:"last_active"`
	City          string `json:"city"`
	StateProvince string `json:"state_prov"`
	ZipCode       string `json:"zip_code"`
	Country       string `json:"country"`
	RookieYear    int    `json:"rookie_year"`
	Website       string `json:"website"`
}

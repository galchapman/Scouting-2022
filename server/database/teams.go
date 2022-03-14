package database

func (db *Database) InsertTeam(team int, name string) (Team, error) {
	_, err := db.db.Exec("INSERT INTO TEAMS (TEAM, NAME) VALUES ($1, $2)", team, name)
	return Team{team, name}, err
}

func (db *Database) GetTeam(team int) (Team, error) {
	row := db.db.QueryRow("SELECT NAME FROM TEAMS WHERE TEAM = $1", team)

	Team := Team{TeamNumber: team}
	err := row.Scan(&Team.Name)
	return Team, err
}

func (db *Database) GetTeamNotes(team Team) (string, error) {
	row := db.db.QueryRow("SELECT NOTES FROM TEAMS WHERE TEAM = $1", team.TeamNumber)

	var notes string
	err := row.Scan(&notes)
	return notes, err
}

func (db *Database) SetTeamNotes(team Team, notes string) error {
	_, err := db.db.Exec("UPDATE TEAMS SET NOTES = $1 WHERE TEAM = $2", notes, team.TeamNumber)
	return err
}

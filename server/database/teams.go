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

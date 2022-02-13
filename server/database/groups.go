package database

func (db *Database) insertGroup(team int, name string) (Team, error) {
	_, err := db.db.Exec("INSERT INTO GROUPS (TEAM, NAME) VALUES ($1, $2)", team, name)
	return Team{team, name}, err
}

package database

import "database/sql"

func (db *Database) InsertGame(game Game) error {
	_, err := db.db.Exec("INSERT INTO GAMES (ID, RED_TEAM1, RED_TEAM2, BLUE_TEAM1, BLUE_TEAM2, RED_SCOUTER1, RED_SCOUTER2, BLUE_SCOUTER1, BLUE_SCOUTER2) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		game.ID, game.Red1.TeamNumber, game.Red2.TeamNumber, game.Blue1.TeamNumber, game.Blue2.TeamNumber,
		game.ScouterRed1.ID, game.ScouterRed2.ID, game.ScouterBlue1.ID, game.ScouterBlue2.ID)
	return err
}

func (db *Database) UpdateGame(game Game) error {
	_, err := db.db.Exec("UPDATE GAMES SET RED_TEAM1 = $1, RED_TEAM2 = $2, BLUE_TEAM1 = $3, BLUE_TEAM2 = $4, RED_SCOUTER1 = $5, RED_SCOUTER2 = $6, BLUE_SCOUTER1 = $7, BLUE_SCOUTER2 = $8 WHERE ID = $9",
		game.Red1.TeamNumber, game.Red2.TeamNumber, game.Blue1.TeamNumber, game.Blue2.TeamNumber,
		game.ScouterRed1.ID, game.ScouterRed2.ID, game.ScouterBlue1.ID, game.ScouterBlue2.ID, game.ID)
	return err
}

func (db *Database) GetGames() ([]Game, error) {
	var games []Game
	var err error
	var rows *sql.Rows

	var Red1 int
	var Red2 int
	var Blue1 int
	var Blue2 int
	var ScoutRed1 int
	var ScoutRed2 int
	var ScoutBlue1 int
	var ScoutBlue2 int

	rows, err = db.db.Query("SELECT ID, GAME_TYPE, RED_TEAM1, RED_TEAM2, BLUE_TEAM1, BLUE_TEAM2, RED_SCOUTER1, RED_SCOUTER2, BLUE_SCOUTER1, BLUE_SCOUTER2 FROM GAMES")

	for rows.Next() {
		var game Game
		err = rows.Scan(&game.ID, &game.GameType, &Red1, &Red2, &Blue1, &Blue2, &ScoutRed1, &ScoutRed2, &ScoutBlue1, &ScoutBlue2)
		if err != nil {
			return nil, err
		}

		game.Red1, _ = db.GetTeam(Red1)
		game.Red2, _ = db.GetTeam(Red2)
		game.Blue1, _ = db.GetTeam(Blue1)
		game.Blue2, _ = db.GetTeam(Blue2)

		game.ScouterRed1, _ = db.GetUser(ScoutRed1)
		game.ScouterRed2, _ = db.GetUser(ScoutRed2)
		game.ScouterBlue1, _ = db.GetUser(ScoutBlue1)
		game.ScouterBlue2, _ = db.GetUser(ScoutBlue2)

		games = append(games, game)
	}

	return games, nil
}

func (db *Database) GetGame(ID int) (Game, error) {
	var game Game

	game.ID = ID

	var Red1 int
	var Red2 int
	var Blue1 int
	var Blue2 int
	var ScoutRed1 int
	var ScoutRed2 int
	var ScoutBlue1 int
	var ScoutBlue2 int

	row := db.db.QueryRow("SELECT RED_TEAM1, RED_TEAM2, BLUE_TEAM1, BLUE_TEAM2, RED_SCOUTER1, RED_SCOUTER2, BLUE_SCOUTER1, BLUE_SCOUTER2 FROM GAMES WHERE ID = $1", ID)

	err := row.Scan(&Red1, &Red2, &Blue1, &Blue2, &ScoutRed1, &ScoutRed2, &ScoutBlue1, &ScoutBlue2)
	if err != nil {
		return Game{}, err
	}

	game.Red1, _ = db.GetTeam(Red1)
	game.Red2, _ = db.GetTeam(Red2)
	game.Blue1, _ = db.GetTeam(Blue1)
	game.Blue2, _ = db.GetTeam(Blue2)

	game.ScouterRed1, _ = db.GetUser(ScoutRed1)
	game.ScouterRed2, _ = db.GetUser(ScoutRed2)
	game.ScouterBlue1, _ = db.GetUser(ScoutBlue1)
	game.ScouterBlue2, _ = db.GetUser(ScoutBlue2)

	return game, nil
}

func (db *Database) GetTeamSchedule(currentGame int, team Team) ([]Game, error) {
	var games []Game
	var err error
	var rows *sql.Rows

	var Red1 int
	var Red2 int
	var Blue1 int
	var Blue2 int

	rows, err = db.db.Query("SELECT ID, RED_TEAM1, RED_TEAM2, BLUE_TEAM1, BLUE_TEAM2 FROM GAMES WHERE ID > $1 AND (RED_TEAM1 = $2 OR RED_TEAM2 = $2 OR BLUE_TEAM1 = $2 OR BLUE_TEAM2 = $2)", currentGame, team.TeamNumber)

	for rows.Next() {
		var game Game
		err = rows.Scan(&game.ID, &Red1, &Red2, &Blue1, &Blue2)
		if err != nil {
			return nil, err
		}

		game.Red1, _ = db.GetTeam(Red1)
		game.Red2, _ = db.GetTeam(Red2)
		game.Blue1, _ = db.GetTeam(Blue1)
		game.Blue2, _ = db.GetTeam(Blue2)

		games = append(games, game)
	}

	return games, nil
}

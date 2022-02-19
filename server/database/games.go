package database

import "database/sql"

func (db *Database) InsertGame(game Game) error {
	_, err := db.db.Exec("INSERT INTO GAMES (ID, RedTeam1, RedTeam2, BlueTeam1, BlueTeam2, RedScouter1, RedScouter2, BlueScouter1, BlueScouter2) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		game.ID, game.Red1, game.Red2, game.Blue1, game.Blue2, game.ScouterRed1, game.ScouterRed2, game.ScouterBlue1, game.ScouterBlue2)
	return err
}

func (db *Database) UpdateGame(game Game) error {
	_, err := db.db.Exec("UPDATE GAMES SET RedTeam1 = $1, RedTeam2 = $2, BlueTeam1 = $3, BlueTeam2 = $4, RedScouter1 = $5, RedScouter2 = $6, BlueScouter1 = $7, BlueScouter2 = $8 WHERE ID = $9",
		game.Red1, game.Red2, game.Blue1, game.Blue2, game.ScouterRed1, game.ScouterRed2, game.ScouterBlue1, game.ScouterBlue2, game.ID)
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

	rows, err = db.db.Query("SELECT ID, RedTeam1, RedTeam2, BlueTeam1, BlueTeam2, RedScouter1, RedScouter2, BlueScouter1, BlueScouter2 FROM GAMES")

	for rows.Next() {
		var game Game
		err = rows.Scan(&game.ID, &Red1, &Red2, &Red2, &Blue1, &Blue2, &ScoutRed1, &ScoutRed2, &ScoutBlue1, &ScoutBlue2)
		if err != nil {
			return nil, err
		}

		game.Red1, _ = db.GetTeam(Red1)
		game.Red2, _ = db.GetTeam(Red2)
		game.Blue1, _ = db.GetTeam(Blue1)
		game.Blue2, _ = db.GetTeam(Blue2)

		game.ScouterRed1, _ = db.GetUser(ScoutRed1)
		game.ScouterRed1, _ = db.GetUser(ScoutRed2)
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

	row := db.db.QueryRow("SELECT RedTeam1, RedTeam2, BlueTeam1, BlueTeam2, RedScouter1, RedScouter2, BlueScouter1, BlueScouter2 FROM GAMES")

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

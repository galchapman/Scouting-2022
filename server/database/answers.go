package database

func (db *Database) InsertAnswer(ScouterID int, TeamNumber int, Game int, answer FormAnswerResponse) error {
	_, err := db.db.Exec("INSERT INTO FORM_ANSWERS (SCOUTER, GAME, TEAM, Location, TEAM_ELEMENT, AUTO_DUCK, AUTO_STORAGE, AUTO_SHIPPING, CUBE_LEVEL, PARKING, STORAGE, SHIPPING_LOW, SHIPPING_MID, SHIPPING_HIGH, SHARED, DUCKS, CAPPING, WORKING, NOTES) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)",
		ScouterID, Game, TeamNumber, answer.Location, answer.TeamElement, answer.AutoDuck, answer.AutoStorage, answer.AutoShipping, answer.CubeLevel, answer.Parking,
		answer.Storage, answer.ShippingLow, answer.ShippingMid, answer.ShippingHigh, answer.Shared, answer.TeleopDucks, answer.Capping, answer.Worked, answer.Notes)
	return err
}

func (db *Database) GetTeamsGames() ([]FormAnswer, error) {
	rows, err := db.db.Query(`SELECT FORM_ANSWERS.ID, SCOUTER, GAME, GAME_TYPE, GAMES.RED_TEAM1, GAMES.RED_TEAM2, GAMES.BLUE_TEAM1,
       GAMES.BLUE_TEAM2, GAMES.RED_SCOUTER1, GAMES.RED_SCOUTER2, GAMES.BLUE_SCOUTER1, GAMES.BLUE_SCOUTER2,
       TEAMS.TEAM, TEAMS.NAME, Location, TEAM_ELEMENT, AUTO_DUCK, AUTO_STORAGE,
       AUTO_SHIPPING, CUBE_LEVEL, PARKING, STORAGE, SHIPPING_LOW, SHIPPING_MID, SHIPPING_HIGH, SHARED, CAPPING, WORKING,
       NOTES FROM FORM_ANSWERS INNER JOIN GAMES ON GAME = GAMES.ID INNER JOIN TEAMS on TEAMS.TEAM = FORM_ANSWERS.TEAM`)

	if err != nil {
		return nil, err
	}

	var answers []FormAnswer
	for rows.Next() {
		var answer FormAnswer
		err = rows.Scan(&answer.ID, &answer.Scouter.ID, &answer.Game.ID, &answer.Game.GameType, &answer.Game.Red1.TeamNumber, &answer.Game.Red2.TeamNumber,
			&answer.Game.Blue1.TeamNumber, &answer.Game.Blue2.TeamNumber, &answer.Game.ScouterRed1.ID, &answer.Game.ScouterRed2.ID,
			&answer.Game.ScouterBlue1.ID, &answer.Game.ScouterBlue2.ID, &answer.Team.TeamNumber, &answer.Team.Name,
			&answer.Location, &answer.TeamElement, &answer.AutoDuck,
			&answer.AutoStorage, &answer.AutoShipping, &answer.CubeLevel, &answer.Parking, &answer.Storage, &answer.ShippingLow,
			&answer.ShippingMid, &answer.ShippingHigh, &answer.Shared, &answer.Capping, &answer.Worked, &answer.Notes)

		if answer.Game.Red1.TeamNumber == answer.Team.TeamNumber || answer.Game.Red2.TeamNumber == answer.Team.TeamNumber {
			answer.Alliance = "Red"
		} else {
			answer.Alliance = "Blue"
		}

		if err != nil {
			return nil, err
		}

		answers = append(answers, answer)
	}

	return answers, err
}

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

func (db *Database) GetTeamGames(team Team) ([]FormAnswer, error) {
	rows, err := db.db.Query(`SELECT FORM_ANSWERS.ID, SCOUTER, GAME, GAME_TYPE, GAMES.RED_TEAM1, GAMES.RED_TEAM2, GAMES.BLUE_TEAM1,
       GAMES.BLUE_TEAM2, GAMES.RED_SCOUTER1, GAMES.RED_SCOUTER2, GAMES.BLUE_SCOUTER1, GAMES.BLUE_SCOUTER2,
       TEAMS.TEAM, TEAMS.NAME, Location, TEAM_ELEMENT, AUTO_DUCK, AUTO_STORAGE,
       AUTO_SHIPPING, CUBE_LEVEL, PARKING, STORAGE, SHIPPING_LOW, SHIPPING_MID, SHIPPING_HIGH, SHARED, CAPPING, WORKING,
       NOTES FROM FORM_ANSWERS INNER JOIN GAMES ON GAME = GAMES.ID INNER JOIN TEAMS on TEAMS.TEAM = FORM_ANSWERS.TEAM WHERE FORM_ANSWERS.TEAM = TEAMS.TEAM`, team.TeamNumber)

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

func (db *Database) InsertSuperFormAnswer(answer SupervisorForm) error {
	_, err := db.db.Exec("INSERT INTO SUPERVISOR_FORM (SCOUTER, GAME, NEAR_RED_INTERFERENCE, NEAR_BLUE_INTERFERENCE, RED1_PENALTY, RED1_NOTES, RED2_PENALTY, RED2_NOTES, BLUE1_PENALTY, BLUE1_NOTES, BLUE2_PENALTY, BLUE2_NOTES) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		answer.Scouter.ID, answer.Game.ID, answer.NearRedInterference, answer.NearBlueInterference,
		answer.Red1Penalty, answer.Red1Notes, answer.Red2Penalty, answer.Red2Notes,
		answer.Blue1Penalty, answer.Blue1Notes, answer.Blue2Penalty, answer.Blue2Notes)
	return err
}

func (db *Database) GetTeamSuperForms(team Team) ([]TeamSupervisorForm, error) {
	rows, err := db.db.Query("SELECT SUPERVISOR_FORM.ID, SUPERVISOR_FORM.SCOUTER, SUPERVISOR_FORM.GAME, RED_TEAM1, RED_TEAM2, BLUE_TEAM1, BLUE_TEAM2, FORM_ANSWERS.Location, NEAR_RED_INTERFERENCE, NEAR_BLUE_INTERFERENCE, RED1_PENALTY, RED1_NOTES, RED2_PENALTY, RED2_NOTES, BLUE1_PENALTY, BLUE1_NOTES, BLUE2_PENALTY, BLUE2_NOTES FROM SUPERVISOR_FORM INNER JOIN GAMES ON SUPERVISOR_FORM.GAME = GAMES.ID INNER JOIN FORM_ANSWERS ON GAMES.ID = FORM_ANSWERS.GAME AND FORM_ANSWERS.TEAM = $1 WHERE GAMES.RED_TEAM1 = $1 OR GAMES.RED_TEAM2 = $1 OR GAMES.BLUE_TEAM1 = $1 OR GAMES.BLUE_TEAM2 = $1",
		team.TeamNumber)
	if err != nil {
		return nil, err
	}

	var forms []TeamSupervisorForm
	for rows.Next() {
		var form SupervisorForm
		var location string
		err = rows.Scan(&form.ID, &form.Scouter.ID, &form.Game.ID, &form.Game.Red1.TeamNumber, &form.Game.Red2.TeamNumber,
			&form.Game.Blue1.TeamNumber, &form.Game.Blue2.TeamNumber, &location, &form.NearRedInterference, &form.NearBlueInterference,
			&form.Red1Penalty, &form.Red1Notes, &form.Red2Penalty, &form.Red2Notes,
			&form.Blue1Penalty, &form.Blue1Notes, &form.Blue2Penalty, &form.Blue2Notes)
		if err != nil {
			return nil, err
		}

		var teamForm TeamSupervisorForm
		switch team.TeamNumber {
		case form.Game.Red1.TeamNumber:
			teamForm.Notes = form.Red1Notes
			teamForm.Penalty = form.Red1Penalty
			break
		case form.Game.Red2.TeamNumber:
			teamForm.Notes = form.Red2Notes
			teamForm.Penalty = form.Red2Penalty
			break
		case form.Game.Blue1.TeamNumber:
			teamForm.Notes = form.Blue1Notes
			teamForm.Penalty = form.Blue1Penalty
			break
		case form.Game.Blue2.TeamNumber:
			teamForm.Notes = form.Blue2Notes
			teamForm.Penalty = form.Blue2Penalty
			break
		}

		if location == "Near" {
			if form.Game.Red1.TeamNumber == team.TeamNumber || form.Game.Red2.TeamNumber == team.TeamNumber {
				teamForm.Interference = form.NearRedInterference
			} else {
				teamForm.Interference = form.NearBlueInterference
			}
		} else {
			teamForm.Interference = false
		}

		forms = append(forms, teamForm)
	}
	return forms, nil
}

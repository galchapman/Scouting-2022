package database

func (db *Database) InsertAnswer(ScouterID int, TeamNumber int, Game int, answer FormAnswerResponse) error {
	_, err := db.db.Exec("INSERT INTO FORM_ANSWERS (SCOUTER, GAME, TEAM, Location, TEAM_ELEMENT, AUTO_DUCK, AUTO_STORAGE, AUTO_SHIPPING, CUBE_LEVEL, PARKING, STORAGE, SHIPPING_LOW, SHIPPING_MID, SHIPPING_HIGH, SHARED, DUCKS, CAPPING, WORKING, NOTES) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)",
		ScouterID, Game, TeamNumber, answer.Location, answer.TeamElement, answer.AutoDuck, answer.AutoStorage, answer.AutoShipping, answer.CubeLevel, answer.Parking,
		answer.Storage, answer.ShippingLow, answer.ShippingMid, answer.ShippingHigh, answer.Shared, answer.TeleopDucks, answer.Capping, answer.Worked, answer.Notes)
	return err
}

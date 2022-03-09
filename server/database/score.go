package database

func (score *TeamScore) Add(answer FormAnswer) {
	score.GamesPlayed++
	// --- Autonomous ---
	var autoScore int
	// Duck
	if answer.AutoDuck == "Yes" {
		autoScore += 10
		if answer.Location == "Near" {
			score.TotalDucksCount++
		}
	}
	// Storage
	storageScore := answer.AutoStorage * 2
	autoScore += storageScore
	// Shipping hub
	shippingScore := answer.AutoShipping * 6
	autoScore += shippingScore
	// Vision
	if answer.Game.GameType != "" {
		score.AddVisionScore(answer.Game.GameType, answer.CubeLevel, answer.TeamElement == "TeamStatue")
	}
	// Parking
	var parkingScore int
	switch answer.Parking {
	case "full warehouse":
		parkingScore = 10
		break
	case "partial warehouse":
		parkingScore = 5
		break
	case "full storage":
		parkingScore = 6
		break
	case "partial storage":
		parkingScore = 3
		break
	}
	autoScore += parkingScore
	// Update scores
	score.TotalScore += autoScore
	score.AutoTotalScore += autoScore
	if answer.Location == "Near" {
		score.AutoStartedNear++
		score.TotalTowerScoreAuto += shippingScore
	}

	// --- Teleop ---
	var teleopScore int
	// Storage
	storageScore = answer.Storage
	teleopScore += storageScore
	// Shipping hub
	shippingScore = answer.ShippingLow*2 + answer.ShippingMid*4 + answer.ShippingHigh*6
	teleopScore += shippingScore
	score.TotalTowerScore += shippingScore
	// Shared hub
	sharedScore := answer.Shared
	teleopScore += sharedScore
	score.TotalSharedScore += sharedScore
	// Update scores
	score.TotalScore += teleopScore

	// --- End Game ---
	var endGameScore int
	// Ducks
	duckScore := answer.TeleopDucks * 6
	score.TotalDucksCount += answer.TeleopDucks
	endGameScore += duckScore
	// Capping
	if answer.Capping == "YesCap" {
		endGameScore += 15
		score.TotalShippingElementsPlaced++
	}

	// --- More Random Stuff ---
	if answer.Notes != "" {
		score.Notes = append(score.Notes, answer.Notes)
	}
	if answer.Worked == "All" || answer.Worked == "AllTeleop" {
		score.Worked++
	}
}

func (score *TeamScore) AddVisionScore(gameType string, cubeLevel int, teamElement bool) {
	var visionScore int
	if teamElement {
		visionScore = 20
	} else {
		visionScore = 10
	}

	if !((gameType == "A" && cubeLevel == 1) || (gameType == "B" && cubeLevel == 2) || (gameType == "C" && cubeLevel == 3)) {
		visionScore = 0
	}

	score.TotalScore += visionScore
	score.AutoTotalScore += visionScore
}

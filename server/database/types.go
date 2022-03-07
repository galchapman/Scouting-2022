package database

const (
	ScouterRole     = iota
	ViewerRole      = iota
	FieldPlayerRole = iota
	ManagerRole     = iota
	AdminRole       = iota
)

type User struct {
	ID             int
	Name           string
	hashedPassword []byte
	ScreenName     string
	Role           int
}

type Team struct {
	TeamNumber int
	Name       string
}

type Game struct {
	ID       int
	GameType string

	Red1  Team
	Red2  Team
	Blue1 Team
	Blue2 Team

	ScouterRed1  User
	ScouterRed2  User
	ScouterBlue1 User
	ScouterBlue2 User
}

type FormAnswer struct {
	ID       int
	Scouter  User
	Team     Team
	Game     Game
	Alliance int
	Location string
	// Auto
	TeamElement  string
	AutoDuck     string
	AutoStorage  int
	AutoShipping int
	CubeLevel    int
	Parking      string
	// Teleop
	Storage      int
	ShippingLow  int
	ShippingMid  int
	ShippingHigh int
	Shared       int
	TeleopDucks  int
	Capping      string
	Worked       string
	Notes        string
}

type FormAnswerResponse struct {
	// Auto
	Location     string
	TeamElement  string
	AutoDuck     string
	AutoStorage  int
	AutoShipping int
	CubeLevel    int
	Parking      string
	// Teleop
	Storage      int
	ShippingLow  int
	ShippingMid  int
	ShippingHigh int
	Shared       int
	TeleopDucks  int
	Capping      string
	Worked       string
	Notes        string
}

type FieldFormAnswer struct {
	ID      int
	Scouter User
	Team    Team

	Cooperation   int
	Communication int
	Kindness      int
	Suitability   int
	Notes         int
}

type TeamScore struct {
	TotalScore  int
	GamesPlayed int
	Worked      int
	Notes       []string

	AutoTotalScore int

	AutoStartedNear     int // Amount of times robot started near the audience
	TotalDucksCountAuto int // Amount of ducks placed when the robot started near the audience
	TotalTowerScoreAuto int // The sum of points scored on the tower during the autonomous stage when the robot started near the audience

	TotalTowerScore             int
	TotalSharedScore            int
	TotalDucksCount             int
	TotalShippingElementsPlaced int
}

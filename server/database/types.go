package database

const (
	ScouterRole     = iota
	ViewerRole      = iota
	FieldPlayerRole = iota
	ManagerRole     = iota
	AdminRole       = iota
)

const (
	RedAlliance  = 0
	BlueAlliance = 1
)

const (
	LocationFar  = 0
	LocationNear = 1
)

// Fake bool
const (
	Succeed     = 0
	Failed      = 1
	HasNotTried = 2
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
	GameType int

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
	AutoStorage  string
	AutoShipping string
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
	AutoStorage  string
	AutoShipping string
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

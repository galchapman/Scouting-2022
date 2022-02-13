package database

const (
	ScouterRole     = iota
	ViewerRole      = iota
	FieldPlayerRole = iota
	Manager         = iota
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
	ID int

	Red1  *Team
	Red2  *Team
	Blue1 *Team
	Blue2 *Team

	ScouterRed1  *User
	ScouterRed2  *User
	ScouterBlue1 *User
	ScouterBlue2 *User
}

type FormAnswer struct {
	ID       int
	Scouter  *User
	Team     *Team
	Game     *Game
	Alliance int
	Location int
	// Auto
	HasTeamElement bool
	AutoDuck       uint8 // Fake bool
	AutoStorage    uint8 // Fake bool
	AutoShipping   uint8 // Fake bool
	AutoShared     uint8 // Fake bool
	CubeLevel      int
	Parking        uint8 // Fake bool
	// Teleop
	Storage      int
	ShippingLow  int
	ShippingMid  int
	ShippingHigh int
	Shard        int
	TeleopDucks  int
	Capping      uint8 // Fake bool
	Worked       string
	Notes        string
}

type FieldFormAnswer struct {
	ID      int
	Scouter *User
	Team    *Team

	Cooperation   int
	Communication int
	Kindness      int
	Suitability   int
	Notes         int
}

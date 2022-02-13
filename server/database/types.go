package database

const (
	ScouterRole     = iota
	ViewerRole      = iota
	FieldPlayerRole = iota
	AdminRole       = iota
	Manager         = iota
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
	ID         int
	Name       string
	Password   []byte
	ScreenName string
	Role       int
}

type Group struct {
	TeamNumber int
	Name       string
}

type Game struct {
	ID int

	Red1  *Group
	Red2  *Group
	Blue1 *Group
	Blue2 *Group

	ScouterRed1  *User
	ScouterRed2  *User
	ScouterBlue1 *User
	ScouterBlue2 *User
}

type FormAnswer struct {
	ID       int
	Scouter  *User
	Team     *Group
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
	Team    *Group

	Cooperation   int
	Communication int
	Kindness      int
	Suitability   int
	Notes         int
}

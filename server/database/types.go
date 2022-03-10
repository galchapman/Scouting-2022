package database

const (
	ScouterRole    = iota
	ViewerRole     = iota
	SupervisorRole = iota
	ManagerRole    = iota
	AdminRole      = iota
)

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	hashedPassword []byte
	ScreenName     string
	Role           int
}

type Team struct {
	TeamNumber int    `json:"number"`
	Name       string `json:"name"`
}

type Game struct {
	ID       int    `json:"id"`
	GameType string `json:"game_type"`

	Red1  Team `json:"red_1"`
	Red2  Team `json:"red_2"`
	Blue1 Team `json:"blue_1"`
	Blue2 Team `json:"blue_2"`

	ScouterRed1  User `json:"scouter_red_1"`
	ScouterRed2  User `json:"scouter_red_2"`
	ScouterBlue1 User `json:"scouter_blue_1"`
	ScouterBlue2 User `json:"scouter_blue_2"`
}

type FormAnswer struct {
	ID       int    `json:"id"`
	Scouter  User   `json:"scouter"`
	Team     Team   `json:"team"`
	Game     Game   `json:"game"`
	Alliance string `json:"alliance"`
	Location string `json:"location"`
	// Auto
	TeamElement  string `json:"team_element"`
	AutoDuck     string `json:"auto_duck"`
	AutoStorage  int    `json:"auto_storage"`
	AutoShipping int    `json:"auto_shipping"`
	CubeLevel    int    `json:"cube_level"`
	Parking      string `json:"parking"`
	// Teleop
	Storage      int    `json:"storage"`
	ShippingLow  int    `json:"shipping_low"`
	ShippingMid  int    `json:"shipping_mid"`
	ShippingHigh int    `json:"shipping_high"`
	Shared       int    `json:"shared"`
	TeleopDucks  int    `json:"teleop_ducks"`
	Capping      string `json:"capping"`
	Worked       string `json:"worked"`
	Notes        string `json:"notes"`
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

type SupervisorForm struct {
	ID                   int
	Scouter              User
	Game                 Game
	NearRedInterference  bool
	NearBlueInterference bool

	Red1Penalty  string
	Red1Notes    string
	Red2Penalty  string
	Red2Notes    string
	Blue1Penalty string
	Blue1Notes   string
	Blue2Penalty string
	Blue2Notes   string
}

type TeamSupervisorForm struct {
	ID           int
	Interference bool
	Penalty      string
	Notes        string
}

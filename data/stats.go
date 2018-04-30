package data

type Server struct {
	GameVersion string `json:"gameVersion"`
	ServerName string `json:"serverName"`
	ServerPort int `json:"serverPort"`
	Port int `json:"port"`
	HostPlayer int `json:"hostPlayer"`
	Game Game `json:"game"`
	Players []Player `json:"players"`
}

type Game struct {
	SprintEnabled bool `json:"sprintEnabled`
	SprintUnlimitedEnabled bool `json:"sprintUnlimitedEnabled`
	MaxPlayers int `json:"maxPlayers`
	MapName string `json:"mapName`
	MapFile string `json:"mapFile`
	Variant string `json:"variant`
	VariantType string `json:"variantType`
	TeamGame bool `json:"teamGame`
	TeamScores []int `json:"teamScores`
}

type Player struct {
	Name string `json:"name"`
	ClientName string `json:"clientName"`
	ServiceTag string `json:"serviceTag"`
	IP string `json:"ip"`
	Team int `json:"team"`
	PlayerIndex int `json:"playerIndex"`
	UID string `json:"uid"`
	PrimaryColor string `json:"primaryColor"`
	PlayerGameStats PlayerGameStats `json:"playerGameStats"`
	PlayerMedals []PlayerMedal `json:"playerMedals"`
	PlayerWeapons []PlayerWeapon `json:"playerWeapons"`
	OtherStats OtherStats `json:"otherStats"`
	playerVersusPlayerKills []int `json:"playerVersusPlayerKills"`
}

type PlayerGameStats struct {
	Score int `json:"score"`
	Kills int `json:"kills"`
	Assists int `json:"assists"`
	Deaths int `json:"deaths"`
	Betrayals int `json:"betrayals"`
	TimeSpentAlive int `json:"timeSpentAlive"`
	Suicides int `json:"suicides"`
	BestStreak int `json:"bestStreak"`
}

type PlayerMedal struct {
	MedalName string `json:"medalName"`
	Count int `json:"count"`
}

type PlayerWeapon struct {
	WeaponName string `json:"weaponName"`
	WeaponIndex int `json:"weaponIndex"`
	Kills int `json:"kills"`
	KilledBy int `json:"killedBy"`
	BetrayalsWith int `json:"betrayalsWith"`
	SuicidesWith int `json:"suicidesWith"`
	HeadshotsWith int `json:"headshotsWith"`
}

type OtherStats struct {
	NemesisIndex int `json:"nemesisIndex"`
	KingsKilled int `json:"kingsKilled"`
	HumansInfected int `json:"humansInfected"`
	ZombiesKilled int `json:"zombiesKilled"`
	TimeInHill int `json:"timeInHill"`
	TimeControllingHill int `json:"timeControllingHill"`
}
package game

type Games map[string]Game

type Game struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
	Deaths     map[string]int `json:"deaths"`
}

type GameStatus struct {
	Game        int
	GameStarted bool
	Killer      string
	Victim      string
	DeathMode   string
}

func NewGameStatus(game int, gameStarted bool, killer string, victim string, DeathMode string) GameStatus {
	return GameStatus{
		Game:        game,
		GameStarted: gameStarted,
		Killer:      killer,
		Victim:      victim,
		DeathMode:   DeathMode,
	}
}

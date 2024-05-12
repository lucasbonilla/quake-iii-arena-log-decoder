package game

type Games map[string]Game

type Game struct {
	TotalKills int
	Players    []string
	Kills      map[string]int
	Deaths     map[string]int
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

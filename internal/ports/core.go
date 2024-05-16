package ports

import "github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/schemas/game"

type Core interface {
	AddNewGame(gameStatus chan<- game.GameStatus, gameID int)
	AddExistingGame(gameStatus chan<- game.GameStatus, initGameCount int, matches []string)
	ProcessPlayerAsVictim(thisGame game.Game, player string) game.Game
	ProcessPlayerAsKiller(thisGame game.Game, player string, worldDeath bool) game.Game
	GenerateJSONFile(games game.Games) error
	GenerateCustomOutput(games game.Games)
}

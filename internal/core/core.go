package core

import (
	"strings"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/schemas/game"
)

type Adapter struct {
	utils ports.Utils
}

func NewAdapter(utilsP ports.Utils) *Adapter {
	return &Adapter{
		utils: utilsP,
	}
}

func (cA *Adapter) AddNewGame(gameStatus chan<- game.GameStatus, gameID int) {
	gameStatus <- game.NewGameStatus(
		gameID,
		true,
		"",
		"",
		"")
}

func (cA *Adapter) AddExistingGame(gameStatus chan game.GameStatus, initGameCount int, matches []string) {
	gameStatus <- game.NewGameStatus(
		initGameCount,
		false,
		strings.TrimLeft(matches[1], " "),
		matches[2],
		matches[3])
}

func (cA *Adapter) ProcessPlayerAsVictim(thisGame game.Game, player string) game.Game {
	if !cA.utils.PlayerExists(thisGame.Players, player) {
		thisGame.Players = append(thisGame.Players, player)
		thisGame.Kills[player] = 0
	}
	return thisGame
}

func (cA *Adapter) ProcessPlayerAsKiller(thisGame game.Game, player string, worldDeath bool) game.Game {
	thisGame.TotalKills++
	if !cA.utils.PlayerExists(thisGame.Players, player) {
		thisGame.Players = append(thisGame.Players, player)
	}
	if worldDeath {
		thisGame.Kills[player]--
		return thisGame
	}
	thisGame.Kills[player]++
	return thisGame
}

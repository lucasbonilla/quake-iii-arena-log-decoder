package core

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/schemas/game"
	"github.com/olekukonko/tablewriter"
)

const (
	FILE_NAME = "qgames.json"
)

type Adapter struct {
	os     ports.Os
	config ports.Config
	utils  ports.Utils
}

func NewAdapter(osP ports.Os, configP ports.Config, utilsP ports.Utils) *Adapter {
	return &Adapter{
		os:     osP,
		config: configP,
		utils:  utilsP,
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

func (cA *Adapter) AddExistingGame(gameStatus chan<- game.GameStatus, initGameCount int, matches []string) {
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

func (cA *Adapter) GenerateJSONFile(games game.Games) error {
	outPath := cA.config.FileOutPath()
	if _, err := cA.os.Stat(outPath); cA.os.IsNotExist(err) {
		if err := cA.os.Mkdir(outPath, 0755); err != nil {
			return err
		}
	}

	fileOut, err := cA.os.Create(outPath + FILE_NAME)

	defer func() {
		cA.os.SetFile(fileOut)
		cA.os.CloseFile()
	}()

	if err != nil {
		return err
	}

	cA.os.SetFile(fileOut)
	encoder := json.NewEncoder(fileOut)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(games)
	if err != nil {
		return err
	}

	return nil
}

func (cA *Adapter) GenerateCustomOutput(games game.Games) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Game ID", "Total Kills", "Players", "Kills", "Deaths"})
	for gameID, game := range games {

		table := tablewriter.NewWriter(os.Stdout)

		table.SetHeader([]string{"Game ID", "Total Kills", "Players", "Kills", "Deaths"})

		players := ""
		for _, player := range game.Players {
			players += player + ", "
		}

		if len(players) > 2 {
			players = players[:len(players)-2]
		}

		kills := ""
		for player, numKills := range game.Kills {
			kills += fmt.Sprintf("%s: %d, ", player, numKills)
		}

		if len(kills) > 2 {
			kills = kills[:len(kills)-2]
		}

		deaths := ""
		for cause, numDeaths := range game.Deaths {
			deaths += fmt.Sprintf("%s: %d, ", cause, numDeaths)
		}

		if len(deaths) > 2 {
			deaths = deaths[:len(deaths)-2]
		}

		table.Append([]string{gameID, fmt.Sprintf("%d", game.TotalKills), players, kills, deaths})

		table.Render()
		fmt.Println()
	}
}

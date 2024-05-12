package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/schemas/game"
)

type Adapter struct {
	os     ports.Os
	utils  ports.Utils
	config ports.Config
	logger ports.Logger
}

func NewAdapter(fileP ports.Os, utilsP ports.Utils, configP ports.Config, loggerP ports.Logger) *Adapter {
	return &Adapter{
		os:     fileP,
		utils:  utilsP,
		config: configP,
		logger: loggerP,
	}
}

func (aA *Adapter) Run() {
	err := aA.os.Open("./logfile/qgames.log")
	if err != nil {
		aA.logger.Error(err.Error())
		return
	}
	defer aA.os.Close()

	// Criar um scanner para ler o arquivo linha por linha
	scanner := aA.os.Scanner()
	aA.os.SetScanner(scanner)

	// Compile a expressão regular para identificar o padrão InitGame
	regexInitGame := regexp.MustCompile(`^\s*\d+:\d+\s+InitGame:`)
	regexDeath := regexp.MustCompile(`(?P<killer>[\w\s<>]+) killed (?P<victim>[\w\s<>]+) by (?P<death_mode>\w+)$`)

	// Variável para contar as ocorrências de InitGame
	var initGameCount int = 0

	gameStatus := make(chan game.GameStatus)

	// Processar cada linha do arquivo
	// var gameStr string
	go func() {
		for aA.os.Scan() {
			line := aA.os.Text()

			if regexInitGame.MatchString(line) {
				initGameCount++
				gameStatus <- game.NewGameStatus(
					initGameCount,
					true,
					"",
					"",
					"")
				continue
			}

			matches := regexDeath.FindStringSubmatch(line)

			if len(matches) > 0 {
				gameStatus <- game.NewGameStatus(
					initGameCount,
					false,
					strings.TrimLeft(matches[1], " "),
					matches[2],
					matches[3])
			}
		}
		close(gameStatus)
	}()
	games := make(game.Games, 0)
	for gs := range gameStatus {
		gameStr := fmt.Sprintf("game_%s", strconv.Itoa(gs.Game))
		if _, ok := games[gameStr]; !ok {
			games[gameStr] = game.Game{
				TotalKills: 0,
				Players:    make([]string, 0),
				Kills:      make(map[string]int),
				Deaths:     make(map[string]int),
			}
		}
		if gs.GameStarted {
			continue
		}
		thisGame := games[gameStr]
		thisGame.Deaths[gs.DeathMode]++
		if gs.Killer == "<world>" {
			thisGame.TotalKills++
			if !aA.utils.PlayerExists(thisGame.Players, gs.Victim) {
				thisGame.Players = append(thisGame.Players, gs.Victim)
			}
			thisGame.Kills[gs.Victim]--
			games[gameStr] = thisGame
			continue
		}
		if !aA.utils.PlayerExists(thisGame.Players, gs.Killer) {
			thisGame.Players = append(thisGame.Players, gs.Killer)
		}
		thisGame.TotalKills++
		thisGame.Kills[gs.Killer]++
		games[gameStr] = thisGame
	}

	fmt.Println(games)
	// fmt.Println(gameStatus)

	// Verificar se houve erros durante a leitura do arquivo
	if err := aA.os.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}

	// Imprimir o número de ocorrências de InitGame
	fmt.Println("Número de ocorrências de InitGame:", initGameCount)
}

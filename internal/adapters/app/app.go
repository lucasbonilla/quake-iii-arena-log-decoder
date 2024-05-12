package app

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"sync"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/schemas/game"
)

type Adapter struct {
	os     ports.Os
	core   ports.Core
	utils  ports.Utils
	config ports.Config
	logger ports.Logger
}

func NewAdapter(fileP ports.Os, coreP ports.Core, utilsP ports.Utils, configP ports.Config, loggerP ports.Logger) *Adapter {
	return &Adapter{
		os:     fileP,
		core:   coreP,
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

	var initGameCount int = 0

	gameStatus := make(chan game.GameStatus)

	var wg sync.WaitGroup

	go aA.produce(gameStatus, regexInitGame, initGameCount, regexDeath)

	games := make(game.Games, 0)
	var mu sync.Mutex
	for range aA.config.GetNumOfWorkers() {
		wg.Add(1)
		go aA.consume(gameStatus, games, &wg, &mu)
	}
	wg.Wait()

	jsonData, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func (aA *Adapter) produce(gameStatus chan game.GameStatus, regexInitGame *regexp.Regexp, initGameCount int, regexDeath *regexp.Regexp) {
	defer close(gameStatus)

	for aA.os.Scan() {
		if err := aA.os.Err(); err != nil {
			fmt.Println("Erro ao ler o arquivo:", err)
		}
		line := aA.os.Text()

		if regexInitGame.MatchString(line) {
			initGameCount++
			aA.core.AddNewGame(gameStatus, initGameCount)
			continue
		}

		matches := regexDeath.FindStringSubmatch(line)
		if len(matches) > 0 {
			aA.core.AddExistingGame(gameStatus, initGameCount, matches)
		}
	}
}

func (aA *Adapter) consume(gameStatus chan game.GameStatus, games game.Games, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for gs := range gameStatus {

		gameStr := fmt.Sprintf("game_%s", strconv.Itoa(gs.Game))
		mu.Lock()
		if _, ok := games[gameStr]; !ok {
			games[gameStr] = game.Game{
				TotalKills: 0,
				Players:    make([]string, 0),
				Kills:      make(map[string]int),
				Deaths:     make(map[string]int),
			}
		}
		mu.Unlock()
		if gs.GameStarted {
			continue
		}
		mu.Lock()
		thisGame := games[gameStr]
		thisGame.Deaths[gs.DeathMode]++

		if gs.Killer == "<world>" {
			games[gameStr] = aA.core.ProcessPlayerAsKiller(thisGame, gs.Victim, true)
			mu.Unlock()
			continue
		}
		thisGame = aA.core.ProcessPlayerAsVictim(thisGame, gs.Victim)
		games[gameStr] = aA.core.ProcessPlayerAsKiller(thisGame, gs.Killer, false)
		mu.Unlock()
	}
}

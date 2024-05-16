package core_test

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"testing"

	externalOs "os"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/config"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os/file"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os/scanner"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/utils"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/core"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/schemas/game"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCustomOutput(t *testing.T) {
	var fileP ports.File
	var scannerP ports.Scanner
	var configP ports.Config
	var osP ports.Os
	var utilsP ports.Utils
	configP = config.NewAdpter("../../test/config_test_dev.toml")
	fileP = file.NewAdapter()
	scannerP = scanner.NewAdapter()
	osP = os.NewAdapter(fileP, scannerP)
	utilsP = utils.NewAdapter()
	var coreP ports.Core = core.NewAdapter(osP, configP, utilsP)
	jsonData := `
	{
	    "game_1": {
	        "total_kills": 0,
	        "players": [],
	        "kills": {},
	        "deaths": {}
	    },
	    "game_2": {
	        "total_kills": 11,
	        "players": [
	            "Isgalamido",
	            "Mocinha"
	        ],
	        "kills": {
	            "Isgalamido": -5,
	            "Mocinha": 0
	        },
	        "deaths": {
	            "MOD_FALLING": 1,
	            "MOD_ROCKET_SPLASH": 3,
	            "MOD_TRIGGER_HURT": 7
	        }
	    },
	   "game_3": {
	        "total_kills": 4,
	        "players": [
	            "Mocinha",
	            "Isgalamido",
	            "Zeh",
	            "Dono da Bola"
	        ],
	        "kills": {
	            "Dono da Bola": -1,
	            "Isgalamido": 1,
	            "Mocinha": 0,
	            "Zeh": -2
	        },
	        "deaths": {
	            "MOD_FALLING": 1,
	            "MOD_ROCKET": 1,
	            "MOD_TRIGGER_HURT": 2
	        }
	    }
	}`

	// Decodifica o JSON para a estrutura Games
	var games game.Games
	err := json.Unmarshal([]byte(jsonData), &games)
	if err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}
	coreP.GenerateCustomOutput(games)
}

func TestGenerateJSONFile(t *testing.T) {
	t.Run("errorCreateFolder", func(t *testing.T) {
		var configP ports.Config
		var osP ports.Os
		var utilsP ports.Utils
		configP = config.NewAdpter("../../test/config_test_dev.toml")

		osP = &os.MockedAdapter{
			StatFn: func(filePath string) (fs.FileInfo, error) {
				return nil, errors.New("an error has occurred")
			},
			IsNotExistFn: func(err error) bool {
				return true
			},
			MkdirFn: func(name string, perm fs.FileMode) error {
				return errors.New("an error has occurred")
			},
		}
		utilsP = utils.NewAdapter()
		var coreP ports.Core = core.NewAdapter(osP, configP, utilsP)
		jsonData := `
		{
			"game_1": {
				"total_kills": 0,
				"players": [],
				"kills": {},
				"deaths": {}
			},
			"game_2": {
				"total_kills": 11,
				"players": [
					"Isgalamido",
					"Mocinha"
				],
				"kills": {
					"Isgalamido": -5,
					"Mocinha": 0
				},
				"deaths": {
					"MOD_FALLING": 1,
					"MOD_ROCKET_SPLASH": 3,
					"MOD_TRIGGER_HURT": 7
				}
			},
		   "game_3": {
				"total_kills": 4,
				"players": [
					"Mocinha",
					"Isgalamido",
					"Zeh",
					"Dono da Bola"
				],
				"kills": {
					"Dono da Bola": -1,
					"Isgalamido": 1,
					"Mocinha": 0,
					"Zeh": -2
				},
				"deaths": {
					"MOD_FALLING": 1,
					"MOD_ROCKET": 1,
					"MOD_TRIGGER_HURT": 2
				}
			}
		}`

		// Decodifica o JSON para a estrutura Games
		var games game.Games
		err := json.Unmarshal([]byte(jsonData), &games)
		if err != nil {
			fmt.Println("Erro ao decodificar o JSON:", err)
			return
		}
		coreP.GenerateJSONFile(games)
	})

	t.Run("errorCreateFile", func(t *testing.T) {
		var configP ports.Config
		var osP ports.Os
		var utilsP ports.Utils
		configP = config.NewAdpter("../../test/config_test_dev.toml")

		osP = &os.MockedAdapter{
			StatFn: func(filePath string) (fs.FileInfo, error) {
				return nil, nil
			},
			IsNotExistFn: func(err error) bool {
				return false
			},
			MkdirFn: func(name string, perm fs.FileMode) error {
				return errors.New("an error has occurred")
			},
			CreateFn: func(fileName string) (*externalOs.File, error) {
				return nil, errors.New("an error has occurred")
			},
			CloseFileFn: func() error {
				return nil
			},
		}
		utilsP = utils.NewAdapter()
		var coreP ports.Core = core.NewAdapter(osP, configP, utilsP)
		jsonData := `
		{
			"game_1": {
				"total_kills": 0,
				"players": [],
				"kills": {},
				"deaths": {}
			},
			"game_2": {
				"total_kills": 11,
				"players": [
					"Isgalamido",
					"Mocinha"
				],
				"kills": {
					"Isgalamido": -5,
					"Mocinha": 0
				},
				"deaths": {
					"MOD_FALLING": 1,
					"MOD_ROCKET_SPLASH": 3,
					"MOD_TRIGGER_HURT": 7
				}
			},
		   "game_3": {
				"total_kills": 4,
				"players": [
					"Mocinha",
					"Isgalamido",
					"Zeh",
					"Dono da Bola"
				],
				"kills": {
					"Dono da Bola": -1,
					"Isgalamido": 1,
					"Mocinha": 0,
					"Zeh": -2
				},
				"deaths": {
					"MOD_FALLING": 1,
					"MOD_ROCKET": 1,
					"MOD_TRIGGER_HURT": 2
				}
			}
		}`

		// Decodifica o JSON para a estrutura Games
		var games game.Games
		err := json.Unmarshal([]byte(jsonData), &games)
		if err != nil {
			fmt.Println("Erro ao decodificar o JSON:", err)
			return
		}
		errGenerate := coreP.GenerateJSONFile(games)
		assert.Error(t, errGenerate)
	})
	t.Run("success", func(t *testing.T) {
		var configP ports.Config
		var osP ports.Os
		var utilsP ports.Utils
		configP = config.NewAdpter("../../test/config_test_dev.toml")

		tempFile, _ := externalOs.Create("tempfile.json")
		defer externalOs.Remove(tempFile.Name())

		osP = &os.MockedAdapter{
			StatFn: func(filePath string) (fs.FileInfo, error) {
				return nil, nil
			},
			IsNotExistFn: func(err error) bool {
				return false
			},
			MkdirFn: func(name string, perm fs.FileMode) error {
				return errors.New("an error has occurred")
			},
			CreateFn: func(fileName string) (*externalOs.File, error) {
				return tempFile, nil
			},
			CloseFileFn: func() error {
				return nil
			},
			SetScannerFn: func(scanner *bufio.Scanner) {},
		}
		utilsP = utils.NewAdapter()
		var coreP ports.Core = core.NewAdapter(osP, configP, utilsP)
		jsonData := `
		{
			"game_1": {
				"total_kills": 0,
				"players": [],
				"kills": {},
				"deaths": {}
			},
			"game_2": {
				"total_kills": 11,
				"players": [
					"Isgalamido",
					"Mocinha"
				],
				"kills": {
					"Isgalamido": -5,
					"Mocinha": 0
				},
				"deaths": {
					"MOD_FALLING": 1,
					"MOD_ROCKET_SPLASH": 3,
					"MOD_TRIGGER_HURT": 7
				}
			},
		   "game_3": {
				"total_kills": 4,
				"players": [
					"Mocinha",
					"Isgalamido",
					"Zeh",
					"Dono da Bola"
				],
				"kills": {
					"Dono da Bola": -1,
					"Isgalamido": 1,
					"Mocinha": 0,
					"Zeh": -2
				},
				"deaths": {
					"MOD_FALLING": 1,
					"MOD_ROCKET": 1,
					"MOD_TRIGGER_HURT": 2
				}
			}
		}`

		// Decodifica o JSON para a estrutura Games
		var games game.Games
		err := json.Unmarshal([]byte(jsonData), &games)
		if err != nil {
			fmt.Println("Erro ao decodificar o JSON:", err)
			return
		}
		errGenerate := coreP.GenerateJSONFile(games)
		assert.Nil(t, errGenerate)
	})

}

func TestProcessPlayerAsKiller(t *testing.T) {
	t.Run("playerDeath", func(t *testing.T) {
		var configP ports.Config
		var osP ports.Os
		var utilsP ports.Utils
		configP = config.NewAdpter("../../test/config_test_dev.toml")

		tempFile, _ := externalOs.Create("tempfile.json")
		defer externalOs.Remove(tempFile.Name())

		osP = &os.MockedAdapter{
			StatFn: func(filePath string) (fs.FileInfo, error) {
				return nil, nil
			},
			IsNotExistFn: func(err error) bool {
				return false
			},
			MkdirFn: func(name string, perm fs.FileMode) error {
				return errors.New("an error has occurred")
			},
			CreateFn: func(fileName string) (*externalOs.File, error) {
				return tempFile, nil
			},
			CloseFileFn: func() error {
				return nil
			},
			SetScannerFn: func(scanner *bufio.Scanner) {},
		}
		utilsP = utils.NewAdapter()

		game := game.Game{
			TotalKills: 10,
			Players:    []string{"Player1", "Player2", "Player3"},
			Kills: map[string]int{
				"Player1": 3,
				"Player2": 4,
				"Player3": 3,
			},
			Deaths: map[string]int{
				"Player1": 2,
				"Player2": 1,
				"Player3": 0,
			},
		}

		var coreP ports.Core = core.NewAdapter(osP, configP, utilsP)
		thisGame := coreP.ProcessPlayerAsKiller(game, "Player1", false)
		game.TotalKills++
		assert.Equal(t, game, thisGame)
	})
	t.Run("worldDeath", func(t *testing.T) {
		var configP ports.Config
		var osP ports.Os
		var utilsP ports.Utils
		configP = config.NewAdpter("../../test/config_test_dev.toml")

		tempFile, _ := externalOs.Create("tempfile.json")
		defer externalOs.Remove(tempFile.Name())

		osP = &os.MockedAdapter{
			StatFn: func(filePath string) (fs.FileInfo, error) {
				return nil, nil
			},
			IsNotExistFn: func(err error) bool {
				return false
			},
			MkdirFn: func(name string, perm fs.FileMode) error {
				return errors.New("an error has occurred")
			},
			CreateFn: func(fileName string) (*externalOs.File, error) {
				return tempFile, nil
			},
			CloseFileFn: func() error {
				return nil
			},
			SetScannerFn: func(scanner *bufio.Scanner) {},
		}
		utilsP = utils.NewAdapter()

		game := game.Game{
			TotalKills: 10,
			Players:    []string{"Player1", "Player2", "Player3"},
			Kills: map[string]int{
				"Player1": 3,
				"Player2": 4,
				"Player3": 3,
			},
			Deaths: map[string]int{
				"Player1": 2,
				"Player2": 1,
				"Player3": 0,
			},
		}

		var coreP ports.Core = core.NewAdapter(osP, configP, utilsP)
		thisGame := coreP.ProcessPlayerAsKiller(game, "Player1", true)
		game.TotalKills++
		assert.Equal(t, game, thisGame)
	})
	t.Run("playerDoesntExists", func(t *testing.T) {
		var configP ports.Config
		var osP ports.Os
		var utilsP ports.Utils
		configP = config.NewAdpter("../../test/config_test_dev.toml")

		tempFile, _ := externalOs.Create("tempfile.json")
		defer externalOs.Remove(tempFile.Name())

		osP = &os.MockedAdapter{
			StatFn: func(filePath string) (fs.FileInfo, error) {
				return nil, nil
			},
			IsNotExistFn: func(err error) bool {
				return false
			},
			MkdirFn: func(name string, perm fs.FileMode) error {
				return errors.New("an error has occurred")
			},
			CreateFn: func(fileName string) (*externalOs.File, error) {
				return tempFile, nil
			},
			CloseFileFn: func() error {
				return nil
			},
			SetScannerFn: func(scanner *bufio.Scanner) {},
		}
		utilsP = utils.NewAdapter()

		game := game.Game{
			TotalKills: 10,
			Players:    []string{"Player1", "Player2", "Player3"},
			Kills: map[string]int{
				"Player1": 3,
				"Player2": 4,
				"Player3": 3,
			},
			Deaths: map[string]int{
				"Player1": 2,
				"Player2": 1,
				"Player3": 0,
			},
		}

		var coreP ports.Core = core.NewAdapter(osP, configP, utilsP)
		thisGame := coreP.ProcessPlayerAsKiller(game, "Player4", true)
		game.TotalKills++
		game.Players = append(game.Players, "Player4")
		assert.Equal(t, game, thisGame)
	})
}

func TestProcessPlayerAsVictim(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var configP ports.Config
		var osP ports.Os
		var utilsP ports.Utils
		configP = config.NewAdpter("../../test/config_test_dev.toml")

		tempFile, _ := externalOs.Create("tempfile.json")
		defer externalOs.Remove(tempFile.Name())

		osP = &os.MockedAdapter{
			StatFn: func(filePath string) (fs.FileInfo, error) {
				return nil, nil
			},
			IsNotExistFn: func(err error) bool {
				return false
			},
			MkdirFn: func(name string, perm fs.FileMode) error {
				return errors.New("an error has occurred")
			},
			CreateFn: func(fileName string) (*externalOs.File, error) {
				return tempFile, nil
			},
			CloseFileFn: func() error {
				return nil
			},
			SetScannerFn: func(scanner *bufio.Scanner) {},
		}
		utilsP = utils.NewAdapter()

		game := game.Game{
			TotalKills: 10,
			Players:    []string{"Player1", "Player2", "Player3"},
			Kills: map[string]int{
				"Player1": 3,
				"Player2": 4,
				"Player3": 3,
			},
			Deaths: map[string]int{
				"Player1": 2,
				"Player2": 1,
				"Player3": 0,
			},
		}

		var coreP ports.Core = core.NewAdapter(osP, configP, utilsP)
		thisGame := coreP.ProcessPlayerAsVictim(game, "Player4")
		game.Players = append(game.Players, "Player4")
		assert.Equal(t, game, thisGame)
	})
}

func TestAddExistingGame(t *testing.T) {
	var configP ports.Config
	var osP ports.Os
	var utilsP ports.Utils
	configP = config.NewAdpter("../../test/config_test_dev.toml")

	tempFile, _ := externalOs.Create("tempfile.json")
	defer externalOs.Remove(tempFile.Name())

	osP = &os.MockedAdapter{}

	utilsP = utils.NewAdapter()

	var coreP ports.Core = core.NewAdapter(osP, configP, utilsP)

	gameStatus := make(chan game.GameStatus, 1)
	expectedInitGameCount := 1
	expectedGameStart := false
	expectedKiller := "Player1"
	expectedVictim := "Player2"
	expectedDeathMode := "BOOM"

	expectedGameStatus := game.NewGameStatus(
		expectedInitGameCount,
		expectedGameStart,
		expectedKiller,
		expectedVictim,
		expectedDeathMode)

	coreP.AddExistingGame(gameStatus, 1, []string{"", expectedKiller, expectedVictim, expectedDeathMode})

	close(gameStatus)
	gameStatusComputed := <-gameStatus

	assert.Equal(t, expectedGameStatus, gameStatusComputed)
}
func TestAddNewGame(t *testing.T) {
	var configP ports.Config
	var osP ports.Os
	var utilsP ports.Utils
	configP = config.NewAdpter("../../test/config_test_dev.toml")

	tempFile, _ := externalOs.Create("tempfile.json")
	defer externalOs.Remove(tempFile.Name())

	osP = &os.MockedAdapter{}

	utilsP = utils.NewAdapter()

	var coreP ports.Core = core.NewAdapter(osP, configP, utilsP)

	gameStatus := make(chan game.GameStatus, 1)
	expectedInitGameCount := 1
	expectedGameStart := true

	expectedGameStatus := game.NewGameStatus(
		expectedInitGameCount,
		expectedGameStart,
		"",
		"",
		"")

	coreP.AddNewGame(gameStatus, 1)

	close(gameStatus)
	gameStatusComputed := <-gameStatus

	assert.Equal(t, expectedGameStatus, gameStatusComputed)
}

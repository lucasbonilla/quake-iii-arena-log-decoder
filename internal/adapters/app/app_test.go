package app_test

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"

	externalOs "os"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/app"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/config"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/logger"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os/file"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os/scanner"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/utils"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/core"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
	"github.com/stretchr/testify/assert"
)

const (
	INIT_GAME        = "0:00 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\dmflags\\0\fraglimit\\20\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\baseq3\\g_needpass\\0"
	PLAYER_KILL_DATA = "2:11 Kill: 2 4 6: Dono da Bola killed Zeh by MOD_ROCKET"
	WORLD_KILL_DATA  = "2:04 Kill: 1022 3 19: <world> killed Isgalamido by MOD_FALLING"
)

func TestNewAdapter(t *testing.T) {
	var configP ports.Config
	var loggerP ports.Logger

	var fileP = &file.MockedAdapter{}
	var scannerP = &scanner.MockedAdapter{}
	var osP ports.Os

	var utilsP ports.Utils

	var coreP ports.Core

	var appP ports.App

	configP = config.NewAdpter("../../../test/config_test_dev.toml")
	loggerP = logger.NewAdapter(configP)

	osP = os.NewAdapter(fileP, scannerP)

	utilsP = utils.NewAdapter()

	coreP = core.NewAdapter(osP, configP, utilsP)

	appP = app.NewAdapter(osP, coreP, utilsP, configP, loggerP)

	assert.NotNil(t, appP)
}

func TestProduce(t *testing.T) {
	t.Run("withNoData", func(t *testing.T) {
		var configP ports.Config
		var loggerP ports.Logger

		var fileP = &file.MockedAdapter{
			OpenFn: func(path string) (*externalOs.File, error) {
				return &externalOs.File{}, nil
			},
			ScannerFn: func() *bufio.Scanner {
				return &bufio.Scanner{}
			},
			CloseFn: func() error {
				return nil
			},
		}
		var scannerP = &scanner.MockedAdapter{}
		var osP ports.Os

		var utilsP ports.Utils

		var coreP ports.Core

		var appP ports.App

		configP = config.NewAdpter("../../../test/config_test_dev.toml")
		loggerP = logger.NewAdapter(configP)

		osP = os.NewAdapter(fileP, scannerP)

		utilsP = utils.NewAdapter()

		coreP = core.NewAdapter(osP, configP, utilsP)

		appP = app.NewAdapter(osP, coreP, utilsP, configP, loggerP)

		appP.Run()

		assert.NotNil(t, appP)
	})
	t.Run("withInitGame", func(t *testing.T) {
		var configP ports.Config
		var loggerP ports.Logger

		var fileP = &file.MockedAdapter{
			OpenFn: func(path string) (*externalOs.File, error) {
				return &externalOs.File{}, nil
			},
			ScannerFn: func() *bufio.Scanner {
				return &bufio.Scanner{}
			},
			CloseFn: func() error {
				return nil
			},
		}
		var scannerP = &scanner.MockedAdapter{
			Lines: []string{INIT_GAME},
			ErrFn: func() error {
				return nil
			},
		}
		var osP ports.Os

		var utilsP ports.Utils

		var coreP ports.Core

		var appP ports.App

		configP = config.NewAdpter("../../../test/config_test_dev.toml")
		loggerP = logger.NewAdapter(configP)

		osP = os.NewAdapter(fileP, scannerP)

		utilsP = utils.NewAdapter()

		coreP = core.NewAdapter(osP, configP, utilsP)

		appP = app.NewAdapter(osP, coreP, utilsP, configP, loggerP)

		appP.Run()

		assert.NotNil(t, appP)
	})

	t.Run("withPlayerData", func(t *testing.T) {
		var configP ports.Config
		var loggerP ports.Logger

		var fileP = &file.MockedAdapter{
			OpenFn: func(path string) (*externalOs.File, error) {
				return &externalOs.File{}, nil
			},
			ScannerFn: func() *bufio.Scanner {
				return &bufio.Scanner{}
			},
			CloseFn: func() error {
				return nil
			},
		}
		var scannerP = &scanner.MockedAdapter{
			Lines: []string{INIT_GAME, PLAYER_KILL_DATA, WORLD_KILL_DATA},
			ErrFn: func() error {
				return nil
			},
		}
		var osP ports.Os

		var utilsP ports.Utils

		var coreP ports.Core

		var appP ports.App

		configP = config.NewAdpter("../../../test/config_test_dev.toml")
		loggerP = logger.NewAdapter(configP)

		osP = os.NewAdapter(fileP, scannerP)

		utilsP = utils.NewAdapter()

		coreP = core.NewAdapter(osP, configP, utilsP)

		appP = app.NewAdapter(osP, coreP, utilsP, configP, loggerP)

		appP.Run()

		assert.NotNil(t, appP)
	})
	t.Run("errorOpenFile", func(t *testing.T) {
		var configP ports.Config
		var loggerP ports.Logger

		var fileP = &file.MockedAdapter{
			OpenFn: func(path string) (*externalOs.File, error) {
				fmt.Println(errors.New("an error has occurred"))
				return nil, errors.New("an error has occurred")
			},
		}
		var scannerP = &scanner.MockedAdapter{}
		var osP ports.Os

		var utilsP ports.Utils

		var coreP ports.Core

		var appP ports.App

		configP = config.NewAdpter("../../../test/config_test_dev.toml")
		loggerP = logger.NewAdapter(configP)

		osP = os.NewAdapter(fileP, scannerP)

		utilsP = utils.NewAdapter()

		coreP = core.NewAdapter(osP, configP, utilsP)

		appP = app.NewAdapter(osP, coreP, utilsP, configP, loggerP)

		old := externalOs.Stdout
		r, w, _ := externalOs.Pipe()
		externalOs.Stdout = w

		appP.Run()

		w.Close()
		externalOs.Stdout = old

		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)

		assert.Equal(t, "an error has occurred\n", buf.String())
	})
	t.Run("errorScann", func(t *testing.T) {
		var configP ports.Config
		var loggerP ports.Logger

		var fileP = &file.MockedAdapter{
			OpenFn: func(path string) (*externalOs.File, error) {
				return &externalOs.File{}, nil
			},
			ScannerFn: func() *bufio.Scanner {
				return &bufio.Scanner{}
			},
			CloseFn: func() error {
				return nil
			},
		}
		var scannerP = &scanner.MockedAdapter{
			Lines: []string{INIT_GAME, PLAYER_KILL_DATA, WORLD_KILL_DATA},
			ErrFn: func() error {
				fmt.Println(errors.New("an error has occurred"))
				return errors.New("an error has occurred")
			},
		}
		var osP ports.Os

		var utilsP ports.Utils

		var coreP ports.Core

		var appP ports.App

		configP = config.NewAdpter("../../../test/config_test_dev.toml")
		loggerP = logger.NewAdapter(configP)

		osP = os.NewAdapter(fileP, scannerP)

		utilsP = utils.NewAdapter()

		coreP = core.NewAdapter(osP, configP, utilsP)

		appP = app.NewAdapter(osP, coreP, utilsP, configP, loggerP)

		old := externalOs.Stdout
		r, w, _ := externalOs.Pipe()
		externalOs.Stdout = w

		appP.Run()

		w.Close()
		externalOs.Stdout = old

		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)

		assert.Equal(t, "an error has occurred\n", buf.String())
	})
}

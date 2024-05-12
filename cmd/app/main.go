package main

import (
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/app"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/config"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/logger"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os/file"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os/scanner"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/utils"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
)

func main() {
	var configP ports.Config
	var loggerP ports.Logger

	var fileP ports.File
	var scannerP ports.Scanner
	var osP ports.Os

	var utilsP ports.Utils

	var appP ports.App

	configP = config.NewAdpter()
	loggerP = logger.NewAdapter(configP)

	fileP = file.NewAdapter()
	scannerP = scanner.NewAdapter()

	osP = os.NewAdapter(fileP, scannerP)

	utilsP = utils.NewAdapter()

	appP = app.NewAdapter(osP, utilsP, configP, loggerP)
	appP.Run()
}

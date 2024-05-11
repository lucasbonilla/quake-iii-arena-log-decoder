package main

import (
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/app"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/config"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/logger"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/os/file"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
)

func main() {
	var configP ports.Config
	var loggerP ports.Logger
	var osP ports.Os
	var fileP ports.File
	var appP ports.App

	configP = config.NewAdpter()
	loggerP = logger.NewAdapter(configP)

	fileP = file.NewAdapter()

	osP = os.NewAdapter(fileP)

	appP = app.NewAdapter(osP, configP, loggerP)
	appP.Run()
}

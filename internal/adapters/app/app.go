package app

import (
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
)

type Adapter struct {
	os     ports.Os
	config ports.Config
	logger ports.Logger
}

func NewAdapter(fileP ports.Os, configP ports.Config, loggerP ports.Logger) *Adapter {
	return &Adapter{
		os:     fileP,
		config: configP,
		logger: loggerP,
	}
}

func (aA *Adapter) Run() {
	var err error
	var filePath string
	filePath, err = aA.config.FilePath()
	if err != nil {
		aA.logger.Error(err.Error())
		return
	}

	err = aA.os.Open(filePath)
	if err != nil {
		aA.logger.Error(err.Error())
		return
	}
	defer aA.os.Close()
}

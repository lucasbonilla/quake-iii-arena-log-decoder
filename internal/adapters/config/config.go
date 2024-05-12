package config

import (
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/schemas/config"
	"github.com/spf13/viper"
)

type Adapter struct {
	APIConfig  *config.API
	fileConfig *config.File
}

func NewAdpter() *Adapter {
	viper.SetDefault("run.workers", "2")

	viper.SetConfigFile("./config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic("it was impossible to read the configuration file")
		}
	}

	apiConfig := config.NewAPIConfig(
		viper.GetString("run.type"),
		viper.GetInt("run.workers"),
	)

	fileConfig := config.NewFileConfig(
		viper.GetString("file.path"))

	return &Adapter{
		APIConfig:  apiConfig,
		fileConfig: fileConfig,
	}
}

func (cA *Adapter) RunType() string {
	return cA.APIConfig.RunType
}

func (cA *Adapter) FilePath() string {
	return cA.fileConfig.Path
}

func (cA *Adapter) GetNumOfWorkers() int {
	return cA.APIConfig.Workers
}

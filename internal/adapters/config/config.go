package config

import (
	"errors"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/schemas/config"
	"github.com/spf13/viper"
)

type Adapter struct {
	APIConfig  *config.API
	fileConfig *config.File
}

func NewAdpter() *Adapter {
	viper.SetConfigFile("./config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic("it was impossible to read the configuration file")
		}
	}

	apiConfig := config.NewAPIConfig(
		viper.GetString("run.type"))

	fileConfig := config.NewFileConfig(
		viper.GetString("file.pathDev"),
		viper.GetString("file.pathPrd"))

	return &Adapter{
		APIConfig:  apiConfig,
		fileConfig: fileConfig,
	}
}

func (cA *Adapter) RunType() string {
	return cA.APIConfig.RunType
}

func (cA *Adapter) FilePath() (string, error) {
	switch cA.RunType() {
	case "debug":
		return cA.fileConfig.PathDEV, nil
	case "prod":
		return cA.fileConfig.PathPRD, nil
	default:
		return "", errors.New("an error occurred while loading the environment variable")
	}
}

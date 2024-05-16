package config

import (
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/schemas/config"
	"github.com/spf13/viper"
)

type Adapter struct {
	APIConfig  *config.API
	fileConfig *config.File
}

func NewAdpter(configFilePath string) *Adapter {
	viper.SetDefault("run.workers", "2")

	viper.SetConfigFile(configFilePath)

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
		viper.GetString("file.pathInDev"),
		viper.GetString("file.pathInProd"),
		viper.GetString("file.pathOutDev"),
		viper.GetString("file.pathOutProd"),
	)

	return &Adapter{
		APIConfig:  apiConfig,
		fileConfig: fileConfig,
	}
}

func (cA *Adapter) RunType() string {
	return cA.APIConfig.RunType
}

func (cA *Adapter) FileInPath() string {
	switch cA.APIConfig.RunType {
	case "prd":
		return cA.fileConfig.PathInProd
	case "dev":
		fallthrough
	default:
		return cA.fileConfig.PathInDev
	}
}

func (cA *Adapter) FileOutPath() string {
	switch cA.APIConfig.RunType {
	case "prd":
		return cA.fileConfig.PathOutProd
	case "dev":
		fallthrough
	default:
		return cA.fileConfig.PathOutDev
	}
}

func (cA *Adapter) GetNumOfWorkers() int {
	return cA.APIConfig.Workers
}

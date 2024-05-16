package config_test

import (
	"errors"
	"testing"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/config"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
	"github.com/stretchr/testify/assert"
)

const (
	expectedPathInDev   = "./dev/path/in/qgames.log"
	expectedPathInProd  = "../../prd/path/in/qgames.log"
	expectedPathOutDev  = "./dev/path/files/out/"
	expectedPathOutProd = "../../prd/path/files/out/"
)

func TestNewAdpter(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var configP ports.Config = config.NewAdpter("../../../test/config_test_dev.toml")
		assert.NotNil(t, configP)
	})
	t.Run("fail", func(t *testing.T) {

		defer func() {
			if r := recover(); r == nil {
				assert.Error(t, errors.New("expected panic"))
			}
		}()
		config.NewAdpter("./fake_config_file.toml")
	})
}

func TestRunType(t *testing.T) {
	var configP ports.Config = config.NewAdpter("../../../test/config_test_prd.toml")
	runType := configP.RunType()
	assert.Equal(t, "prd", runType)
}

func TestFileInPath(t *testing.T) {
	t.Run("prd", func(t *testing.T) {
		var configP ports.Config = config.NewAdpter("../../../test/config_test_prd.toml")
		fileIn := configP.FileInPath()
		assert.Equal(t, expectedPathInProd, fileIn)
	})
	t.Run("dev", func(t *testing.T) {
		var configP ports.Config = config.NewAdpter("../../../test/config_test_dev.toml")
		fileIn := configP.FileInPath()
		assert.Equal(t, expectedPathInDev, fileIn)
	})
}

func TestFileOutPath(t *testing.T) {
	t.Run("prd", func(t *testing.T) {
		var configP ports.Config = config.NewAdpter("../../../test/config_test_prd.toml")
		fileOut := configP.FileOutPath()
		assert.Equal(t, expectedPathOutProd, fileOut)
	})
	t.Run("dev", func(t *testing.T) {
		var configP ports.Config = config.NewAdpter("../../../test/config_test_dev.toml")
		fileOut := configP.FileOutPath()
		assert.Equal(t, expectedPathOutDev, fileOut)
	})
}

func TestGetNumOfWorkers(t *testing.T) {
	var configP ports.Config = config.NewAdpter("../../../test/config_test_prd.toml")
	workers := configP.GetNumOfWorkers()
	assert.Equal(t, 5, workers)
}

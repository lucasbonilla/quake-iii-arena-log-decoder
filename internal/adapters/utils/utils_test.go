package utils_test

import (
	"testing"

	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/adapters/utils"
	"github.com/lucasbonilla/quake-iii-arena-log-decoder/internal/ports"
	"github.com/stretchr/testify/assert"
)

func TestPlayerExists(t *testing.T) {
	t.Run("playerExists", func(t *testing.T) {
		players := []string{"Hank", "Eric", "Diana", "Sheila", "Presto", "Bobby", "Uni"}
		var utilsP ports.Utils = utils.NewAdapter()

		itExists := utilsP.PlayerExists(players, "Sheila")
		assert.True(t, itExists)
	})
	t.Run("playerDoesntExists", func(t *testing.T) {
		players := []string{"Hank", "Eric", "Diana", "Sheila", "Presto", "Bobby", "Uni"}
		var utilsP ports.Utils = utils.NewAdapter()

		itExists := utilsP.PlayerExists(players, "Mestre dos Magos")
		assert.False(t, itExists)
	})
}

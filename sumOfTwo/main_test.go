package sumOfTwo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneral(t *testing.T) {
	t.Run("general run", func(t *testing.T) {
		got := sum(4, 5)
		require.Equal(t, 9, got)
	})
}

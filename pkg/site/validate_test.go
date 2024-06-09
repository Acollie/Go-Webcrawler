package site

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("Test valid string", func(t *testing.T) {
		depth := CheckPageDepth("http://example.com", 10)
		require.Equal(t, depth, true)
	})
	t.Run("Test invalid string", func(t *testing.T) {
		depth := CheckPageDepth("http://example.com/foo/bar/foo", 2)
		require.Equal(t, depth, false)
	})

}

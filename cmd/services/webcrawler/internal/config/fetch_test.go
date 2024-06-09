package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIgnoreList(t *testing.T) {

	t.Run("FetchDDB ignore list ", func(t *testing.T) {
		conf := IgnoreList{
			Website: map[string]bool{
				"github.com": true,
			},
		}
		url := "https://github.com"
		require.Equal(t, conf.Ignore(url), true)
		require.Equal(t, conf.Ignore("https://alexcollie.com"), false)
	})
}

package queue

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewMessage(t *testing.T) {
	message := NewMessage("test")
	require.Equal(t, message.Url, "test")
}

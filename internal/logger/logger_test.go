package logger

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

type loggerConfig struct {
	Level string
}

func (l *loggerConfig) GetLevel() string {
	return l.Level
}

func TestLogger(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		New(&loggerConfig{
			Level: "info",
		})
		require.Equal(t, zerolog.InfoLevel, zerolog.GlobalLevel())

		New(&loggerConfig{
			Level: "error",
		})
		require.Equal(t, zerolog.ErrorLevel, zerolog.GlobalLevel())

		New(&loggerConfig{
			Level: "warn",
		})
		require.Equal(t, zerolog.WarnLevel, zerolog.GlobalLevel())
	})
}

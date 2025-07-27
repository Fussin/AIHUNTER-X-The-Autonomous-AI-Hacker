package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/user/aihunter-x/internal/config"
)

// New initializes the logger.
func New(cfg config.LogConfig) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	level, err := zerolog.ParseLevel(strings.ToLower(cfg.Level))
	if err != nil {
		log.Warn().Msgf("unknown log level: %s. defaulting to 'info'", cfg.Level)
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	if cfg.Format == "json" {
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

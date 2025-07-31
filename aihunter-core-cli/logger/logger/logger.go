package logger

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
	"gopkg.in/natefinch/lumberjack.v2"
)

// New initializes the logger.
// It configures the logger based on the provided configuration.
func New(cfg config.LogConfig, verbose bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	level, err := zerolog.ParseLevel(strings.ToLower(cfg.Level))
	if err != nil {
		log.Warn().Msgf("unknown log level: %s. defaulting to 'info'", cfg.Level)
		level = zerolog.InfoLevel
	}

	if verbose {
		level = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(level)

	var writers []io.Writer
	if cfg.File != "" {
		r := &lumberjack.Logger{
			Filename:   cfg.File,
			MaxSize:    cfg.Rotate.MaxSize,
			MaxBackups: cfg.Rotate.MaxBackups,
			MaxAge:     cfg.Rotate.MaxAge,
			Compress:   cfg.Rotate.Compress,
		}
		writers = append(writers, r)
	}

	if cfg.SessionID == "" {
		cfg.SessionID = uuid.New().String()
	}

	if cfg.Format == "json" {
		writers = append(writers, os.Stdout)
		log.Logger = zerolog.New(io.MultiWriter(writers...)).With().Timestamp().Str("session_id", cfg.SessionID).Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}
		consoleWriter.FormatErrFieldName = func(i interface{}) string {
			return fmt.Sprintf("\x1b[31m%s:\x1b[0m", i)
		}
		consoleWriter.FormatErrFieldValue = func(i interface{}) string {
			return fmt.Sprintf("\x1b[31m%s\x1b[0m", i)
		}
		writers = append(writers, consoleWriter)
		log.Logger = log.Output(io.MultiWriter(writers...)).With().Str("session_id", cfg.SessionID).Logger()
	}

	log.Info().
		Str("version", "0.0.1").
		Interface("config", viper.AllSettings()).
		Msg("Starting AIHUNTER-X")
}

package logger_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
	"github.com/user/aihunter-x/aihunter-core-cli/logger/logger"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	testLogger := zerolog.New(&buf)

	cfg := config.LogConfig{
		Level:  "debug",
		Format: "console",
	}
	logger.New(cfg, true)

	testLogger.Debug().Msg("test message")

	assert.Contains(t, buf.String(), "test message")
}

func TestNew_File(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "log-*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	cfg := config.LogConfig{
		Level:  "debug",
		Format: "json",
		File:   tmpfile.Name(),
	}
	logger.New(cfg, true)

	log.Debug().Msg("test message")

	// The logger is asynchronous, so we need to wait a bit for the log to be written.
	time.Sleep(100 * time.Millisecond)

	content, err := os.ReadFile(tmpfile.Name())
	assert.NoError(t, err)
	assert.Contains(t, string(content), "test message")
}

func TestNew_SessionID(t *testing.T) {
	var buf bytes.Buffer
	testLogger := zerolog.New(&buf)

	cfg := config.LogConfig{
		Level:     "debug",
		Format:    "json",
		SessionID: "test-session-id",
	}
	logger.New(cfg, true)

	log.Debug().Msg("test message")

	assert.Contains(t, buf.String(), "test-session-id")
}

func TestNew_Summary(t *testing.T) {
	var buf bytes.Buffer
	log.Logger = zerolog.New(&buf)

	viper.Set("test-key", "test-value")
	cfg := config.LogConfig{
		Level:  "info",
		Format: "json",
	}
	logger.New(cfg, false)

	assert.Contains(t, buf.String(), "Starting AIHUNTER-X")
	assert.Contains(t, buf.String(), "test-key")
	assert.Contains(t, buf.String(), "test-value")
}

func TestNew_Error(t *testing.T) {
	var outBuf, errBuf bytes.Buffer
	log.Logger = zerolog.New(&outBuf)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cfg := config.LogConfig{
		Level:  "info",
		Format: "console",
	}
	logger.New(cfg, false)

	log.Error().Msg("test error")

	assert.NotContains(t, outBuf.String(), "test error")
	assert.Contains(t, errBuf.String(), "test error")
}

func TestNew_FileRotation(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "log-*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	cfg := config.LogConfig{
		Level:  "debug",
		Format: "json",
		File:   tmpfile.Name(),
		Rotate: config.RotateConfig{
			MaxSize:    1, // 1 MB
			MaxBackups: 1,
			MaxAge:     1,
			Compress:   false,
		},
	}
	logger.New(cfg, true)

	// Write more than 1 MB to trigger rotation
	for i := 0; i < 100000; i++ {
		log.Debug().Msg("test message")
	}

	// The logger is asynchronous, so we need to wait a bit for the log to be written.
	time.Sleep(100 * time.Millisecond)

	// Check that there is a backup file
	files, err := os.ReadDir(filepath.Dir(tmpfile.Name()))
	assert.NoError(t, err)
	assert.Greater(t, len(files), 1)
}

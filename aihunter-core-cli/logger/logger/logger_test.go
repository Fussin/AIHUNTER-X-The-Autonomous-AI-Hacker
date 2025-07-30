package logger

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
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

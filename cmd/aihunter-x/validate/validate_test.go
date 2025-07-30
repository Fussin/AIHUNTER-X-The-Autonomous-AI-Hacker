package validate

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestNewValidateCmd(t *testing.T) {
	var logBuf bytes.Buffer
	log.Logger = zerolog.New(&logBuf)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cmd := NewValidateCmd()
	cmd.Execute()
	assert.Contains(t, logBuf.String(), "validate called")
}

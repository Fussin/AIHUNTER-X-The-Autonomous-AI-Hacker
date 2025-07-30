package recon

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestNewReconCmd(t *testing.T) {
	var logBuf bytes.Buffer
	log.Logger = zerolog.New(&logBuf)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cmd := NewReconCmd()
	cmd.Execute()
	assert.Contains(t, logBuf.String(), "recon called")
}

package scan

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestNewScanCmd(t *testing.T) {
	var outBuf, errBuf bytes.Buffer
	cmd := NewScanCmd()
	cmd.SetOut(&outBuf)
	cmd.SetErr(&errBuf)
	cmd.Flags().BoolP("verbose", "v", false, "verbose output")
	cmd.SetArgs([]string{"-v"})

	var logBuf bytes.Buffer
	log.Logger = zerolog.New(&logBuf)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	cmd.Execute()
	assert.Contains(t, logBuf.String(), "scan called")
}

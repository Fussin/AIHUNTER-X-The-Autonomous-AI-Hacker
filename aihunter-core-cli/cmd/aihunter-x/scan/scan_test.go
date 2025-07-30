package scan

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/user/aihunter-x/aihunter-core-cli/cmd/aihunter-x/scan"
)

func TestNewScanCmd(t *testing.T) {
	var logBuf bytes.Buffer
	log.Logger = zerolog.New(&logBuf)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	cmd := scan.NewScanCmd()
	cmd.SetArgs([]string{"test"})
	cmd.Execute()
	assert.Contains(t, logBuf.String(), "scan called for target: test")
}

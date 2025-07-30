package recon_test

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/user/aihunter-x/aihunter-core-cli/cmd/aihunter-x/recon"
)

func TestNewReconCmd(t *testing.T) {
	var logBuf bytes.Buffer
	log.Logger = zerolog.New(&logBuf)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	p := &recon.Plugin{}
	cmd := p.Commands()[0]
	cmd.SetArgs([]string{"test"})
	cmd.Execute()
	assert.Contains(t, logBuf.String(), "recon called for target: test")
}

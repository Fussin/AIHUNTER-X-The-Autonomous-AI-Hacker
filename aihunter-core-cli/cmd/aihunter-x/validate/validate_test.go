package validate_test

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/user/aihunter-x/aihunter-core-cli/cmd/aihunter-x/validate"
)

func TestNewValidateCmd(t *testing.T) {
	var logBuf bytes.Buffer
	log.Logger = zerolog.New(&logBuf)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	p := &validate.Plugin{}
	cmd := p.Commands()[0]
	cmd.SetArgs([]string{"test"})
	cmd.Execute()
	assert.Contains(t, logBuf.String(), "validate called for finding: test")
}

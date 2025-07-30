package scan_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/user/aihunter-x/aihunter-core-cli/cmd/aihunter-x/scan"
)

func TestNewScanCmd(t *testing.T) {
	var outBuf bytes.Buffer
	p := &scan.Plugin{}
	cmd := p.Commands()[0]
	cmd.SetOut(&outBuf)
	cmd.SetArgs([]string{"test"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, outBuf.String(), "map[target:test vulnerability:xss]")
}

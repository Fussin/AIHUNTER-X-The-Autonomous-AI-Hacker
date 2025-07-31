package utils

import (
	"bytes"
	"errors"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestPrintError(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf
	color.NoColor = false

	err := errors.New("test error")
	PrintError(err)

	assert.Equal(t, "\x1b[31mError: test error\n\x1b[0m", buf.String())
}

func TestPrintWarning(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf
	color.NoColor = false

	msg := "test warning"
	PrintWarning(msg)

	assert.Equal(t, "\x1b[33mWarning: test warning\n\x1b[0m", buf.String())
}

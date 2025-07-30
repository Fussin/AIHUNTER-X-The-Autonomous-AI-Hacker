package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONFormatter(t *testing.T) {
	formatter, err := NewFormatter("json")
	assert.NoError(t, err)

	data := map[string]string{"foo": "bar"}
	formattedData, err := formatter.Format(data)
	assert.NoError(t, err)

	expected := `{
  "foo": "bar"
}`
	assert.Equal(t, expected, string(formattedData))
}

func TestNewFormatter_Unknown(t *testing.T) {
	_, err := NewFormatter("unknown")
	assert.Error(t, err)
}

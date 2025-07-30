package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
)

func TestLoad(t *testing.T) {
	// Create a dummy config file
	content := []byte(`
log:
  level: "debug"
  format: "json"
`)
	tmpfile, err := os.CreateTemp("", "config-*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write(content)
	assert.NoError(t, err)
	err = tmpfile.Close()
	assert.NoError(t, err)

	// Test loading the config file
	cfg, err := config.Load(tmpfile.Name())
	assert.NoError(t, err)
	assert.Equal(t, "debug", cfg.Log.Level)
	assert.Equal(t, "json", cfg.Log.Format)
}

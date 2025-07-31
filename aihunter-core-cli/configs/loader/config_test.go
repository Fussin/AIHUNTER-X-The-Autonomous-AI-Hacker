package config_test

import (
	"os"
	"testing"

	"github.com/spf13/viper"
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

func TestLoad_Defaults(t *testing.T) {
	viper.Reset()
	// Test loading with no config file
	cfg, err := config.Load("")
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "info", cfg.Log.Level)
	assert.Equal(t, "console", cfg.Log.Format)
}

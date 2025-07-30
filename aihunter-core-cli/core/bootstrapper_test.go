package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
	"github.com/user/aihunter-x/aihunter-core-cli/engine"
)

func TestBootstrap(t *testing.T) {
	cfg := &config.Config{}

	// Test with a valid command
	e, err := Bootstrap("scan", cfg)
	assert.NoError(t, err)
	assert.IsType(t, &engine.ScanEngine{}, e)

	// Test with another valid command
	e, err = Bootstrap("recon", cfg)
	assert.NoError(t, err)
	assert.IsType(t, &engine.ReconEngine{}, e)

	// Test with yet another valid command
	e, err = Bootstrap("validate", cfg)
	assert.NoError(t, err)
	assert.IsType(t, &engine.ValidateEngine{}, e)

	// Test with an invalid command
	_, err = Bootstrap("unknown", cfg)
	assert.Error(t, err)
}

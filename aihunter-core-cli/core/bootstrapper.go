package core

import (
	"fmt"

	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
	"github.com/user/aihunter-x/aihunter-core-cli/engine"
)

// Bootstrap bootstraps the engine based on the command name.
func Bootstrap(cmdName string, cfg *config.Config) (Engine, error) {
	switch cmdName {
	case "scan":
		return &engine.ScanEngine{}, nil
	case "recon":
		return &engine.ReconEngine{}, nil
	case "validate":
		return &engine.ValidateEngine{}, nil
	default:
		return nil, fmt.Errorf("unknown command: %s", cmdName)
	}
}

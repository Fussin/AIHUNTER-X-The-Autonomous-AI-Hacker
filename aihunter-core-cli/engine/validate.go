package engine

import (
	"fmt"

	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
)

// ValidateEngine is a dummy implementation of the validate engine.
type ValidateEngine struct{}

// Run runs the validate engine.
func (e *ValidateEngine) Run(cfg *config.Config) error {
	fmt.Println("Running validate engine...")
	return nil
}

package engine

import (
	"fmt"

	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
)

// ReconEngine is a dummy implementation of the recon engine.
type ReconEngine struct{}

// Run runs the recon engine.
func (e *ReconEngine) Run(cfg *config.Config) error {
	fmt.Println("Running recon engine...")
	return nil
}

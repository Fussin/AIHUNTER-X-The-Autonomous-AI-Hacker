package engine

import (
	"fmt"

	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
)

// ScanEngine is a dummy implementation of the scan engine.
type ScanEngine struct{}

// Run runs the scan engine.
func (e *ScanEngine) Run(cfg *config.Config) error {
	fmt.Println("Running scan engine...")
	return nil
}

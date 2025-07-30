package core

import "github.com/user/aihunter-x/aihunter-core-cli/configs/loader"

// Engine is the interface that all engines must implement.
type Engine interface {
	// Run starts the engine.
	// It receives the configuration and returns an error if something goes wrong.
	Run(cfg *config.Config) error
}

// ScanEngine is the interface for the scan engine.
// It embeds the Engine interface.
type ScanEngine interface {
	Engine
}

// ReconEngine is the interface for the recon engine.
// It embeds the Engine interface.
type ReconEngine interface {
	Engine
}

// ValidateEngine is the interface for the validate engine.
// It embeds the Engine interface.
type ValidateEngine interface {
	Engine
}

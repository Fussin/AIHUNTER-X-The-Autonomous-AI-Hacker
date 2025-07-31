package dispatcher

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Plugin is the interface that all plugins must implement.
type Plugin interface {
	// Commands returns a list of commands that the plugin provides.
	Commands() []*cobra.Command
	// PreInit is called before the command is executed.
	PreInit() error
	// OnExit is called when the application is about to exit.
	OnExit() error
}

var plugins []Plugin

// Register registers a new plugin.
// This function should be called by each plugin in its init() function.
func Register(p Plugin) {
	plugins = append(plugins, p)
}

// ResetPlugins resets the list of registered plugins.
// This function is intended for testing purposes only.
func ResetPlugins() {
	plugins = []Plugin{}
}

// AddCommands adds all registered plugin commands to the root command.
func AddCommands(root *cobra.Command) {
	for _, p := range plugins {
		for _, cmd := range p.Commands() {
			// Wrap the RunE function to recover from panics
			if cmd.RunE != nil {
				runE := cmd.RunE
				cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
					defer func() {
						if r := recover(); r != nil {
							err = fmt.Errorf("panic: %v", r)
						}
					}()
					return runE(cmd, args)
				}
			}
			root.AddCommand(cmd)
		}
	}
}

// GetPlugins returns all registered plugins.
func GetPlugins() []Plugin {
	return plugins
}

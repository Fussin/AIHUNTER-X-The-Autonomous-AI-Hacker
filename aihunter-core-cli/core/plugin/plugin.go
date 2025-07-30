package plugin

import "github.com/spf13/cobra"

// Plugin is an interface for plugins.
type Plugin interface {
	// Commands returns a list of commands to be added to the root command.
	Commands() []*cobra.Command
}

var plugins []Plugin

// Register registers a new plugin.
// This function is called by plugins to register themselves.
func Register(p Plugin) {
	plugins = append(plugins, p)
}

// AddCommands adds all registered commands to the root command.
// This function is called by the main application to add all registered commands.
func AddCommands(root *cobra.Command) {
	for _, p := range plugins {
		for _, cmd := range p.Commands() {
			root.AddCommand(cmd)
		}
	}
}

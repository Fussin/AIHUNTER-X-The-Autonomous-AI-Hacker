package dispatcher

import "github.com/spf13/cobra"

// Plugin is the interface that all plugins must implement.
type Plugin interface {
	// Commands returns a list of commands that the plugin provides.
	Commands() []*cobra.Command
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
			root.AddCommand(cmd)
		}
	}
}

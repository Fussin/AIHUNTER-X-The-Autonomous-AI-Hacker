package plugin

import "github.com/spf13/cobra"

// Command is a function that returns a cobra.Command.
type Command func() *cobra.Command

var commands []Command

// Register registers a new command.
// This function is called by plugins to register their commands.
func Register(cmd Command) {
	commands = append(commands, cmd)
}

// AddCommands adds all registered commands to the root command.
// This function is called by the main application to add all registered commands.
func AddCommands(root *cobra.Command) {
	for _, cmd := range commands {
		root.AddCommand(cmd())
	}
}

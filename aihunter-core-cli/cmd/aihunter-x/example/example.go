package example

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Plugin is the example plugin.
type Plugin struct{}

// Commands returns the commands for the example plugin.
func (p *Plugin) Commands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example plugin.",
		Long:  `An example plugin.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from the example plugin!")
		},
	}
	return []*cobra.Command{cmd}
}

// PreInit is called before the command is executed.
func (p *Plugin) PreInit() error {
	fmt.Println("PreInit called from the example plugin!")
	return nil
}

// OnExit is called when the application is about to exit.
func (p *Plugin) OnExit() error {
	fmt.Println("OnExit called from the example plugin!")
	return nil
}

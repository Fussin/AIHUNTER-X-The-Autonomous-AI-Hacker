package recon

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Plugin is the recon plugin.
type Plugin struct{}

// Commands returns the commands for the recon plugin.
func (p *Plugin) Commands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:     "recon [target]",
		Aliases: []string{"r"},
		Short:   "Perform reconnaissance on a target.",
		Long:    `Perform reconnaissance on a target. Provide the target as an argument.`,
		Example: `  aihunter-x recon example.com`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().Msgf("recon called for target: %s", args[0])
		},
	}
	// Add flags here
	return []*cobra.Command{cmd}
}

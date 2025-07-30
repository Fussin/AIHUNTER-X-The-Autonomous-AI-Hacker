package validate

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Plugin is the validate plugin.
type Plugin struct{}

// Commands returns the commands for the validate plugin.
func (p *Plugin) Commands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:     "validate [finding]",
		Aliases: []string{"v"},
		Short:   "Validate a vulnerability finding.",
		Long:    `Validate a vulnerability finding. Provide the finding as an argument.`,
		Example: `  aihunter-x validate "xss in example.com"`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().Msgf("validate called for finding: %s", args[0])
		},
	}
	// Add flags here
	return []*cobra.Command{cmd}
}

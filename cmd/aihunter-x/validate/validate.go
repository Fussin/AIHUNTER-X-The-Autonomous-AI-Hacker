package validate

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewValidateCmd creates a new validate command
func NewValidateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate [finding]",
		Short: "Validate a vulnerability finding.",
		Long:  `Validate a vulnerability finding. Provide the finding as an argument.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().Msgf("validate called for finding: %s", args[0])
		},
	}
	// Add flags here
	return cmd
}

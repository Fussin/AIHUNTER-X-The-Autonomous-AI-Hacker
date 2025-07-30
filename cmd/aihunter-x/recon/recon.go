package recon

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewReconCmd creates a new recon command
func NewReconCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recon [target]",
		Short: "Perform reconnaissance on a target.",
		Long:  `Perform reconnaissance on a target. Provide the target as an argument.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().Msgf("recon called for target: %s", args[0])
		},
	}
	// Add flags here
	return cmd
}

package scan

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewScanCmd creates a new scan command
func NewScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan [target]",
		Short: "Scan a target for vulnerabilities.",
		Long:  `Scan a target for vulnerabilities. Provide the target as an argument.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.Debug().Msgf("scan called for target: %s", args[0])
		},
	}
	// Add flags here
	return cmd
}

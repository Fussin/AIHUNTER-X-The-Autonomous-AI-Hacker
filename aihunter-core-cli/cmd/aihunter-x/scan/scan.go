package scan

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/user/aihunter-x/aihunter-core-cli/logger"
)

// Plugin is the scan plugin.
type Plugin struct{}

// Commands returns the commands for the scan plugin.
func (p *Plugin) Commands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan [target]",
		Short: "Scan a target for vulnerabilities.",
		Long:  `Scan a target for vulnerabilities. Provide the target as an argument.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			output, _ := cmd.Flags().GetString("output")
			formatter, err := logger.NewFormatter(output)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to create formatter")
			}

			// Dummy data for now
			data := map[string]string{"target": args[0], "vulnerability": "xss"}
			formattedData, err := formatter.Format(data)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to format data")
			}
			fmt.Println(string(formattedData))
		},
	}
	cmd.Flags().StringP("output", "o", "console", "output format (console, json)")
	return []*cobra.Command{cmd}
}

package recon

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
	"github.com/user/aihunter-x/aihunter-core-cli/core"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(viper.ConfigFileUsed())
			if err != nil {
				return err
			}

			engine, err := core.Bootstrap(cmd.Name(), cfg)
			if err != nil {
				return err
			}

			return engine.Run(cfg)
		},
	}
	// Add flags here
	return []*cobra.Command{cmd}
}

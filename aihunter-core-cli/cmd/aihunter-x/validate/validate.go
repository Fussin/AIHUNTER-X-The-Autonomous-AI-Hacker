package validate

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
	"github.com/user/aihunter-x/aihunter-core-cli/core"
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

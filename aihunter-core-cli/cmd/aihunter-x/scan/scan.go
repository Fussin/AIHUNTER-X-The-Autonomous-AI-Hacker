package scan

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
	"github.com/user/aihunter-x/aihunter-core-cli/core"
)

// Plugin is the scan plugin.
type Plugin struct{}

// Commands returns the commands for the scan plugin.
func (p *Plugin) Commands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:     "scan [target]",
		Aliases: []string{"s"},
		Short:   "Scan a target for vulnerabilities.",
		Long:    `Scan a target for vulnerabilities. Provide the target as an argument.`,
		Example: `  aihunter-x scan example.com
  aihunter-x scan example.com -o result.json`,
		Args:             cobra.ExactArgs(1),
		PersistentPreRunE: core.InitConfig,
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
	cmd.Flags().StringP("output", "o", "console", "output format (console, json)")
	return []*cobra.Command{cmd}
}

// PreInit is called before the command is executed.
func (p *Plugin) PreInit() error {
	return nil
}

// OnExit is called when the application is about to exit.
func (p *Plugin) OnExit() error {
	return nil
}

package scan

import (
	"fmt"
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
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			shell, _ := cmd.Flags().GetBool("shell")
			if shell {
				fmt.Println("export AIHUNTER_X_TARGET=example.com")
				fmt.Println("export AIHUNTER_X_VULNERABILITY=xss")
				return nil
			}

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
	cmd.Flags().Bool("shell", false, "output results as environment variables")
	cmd.SetHelpFunc(func(cmd *cobra.Command, a []string) {
		output, _ := cmd.Flags().GetString("output")
		if output == "json" {
			cmd.Example = `  aihunter-x scan example.com -o result.json`
		} else {
			cmd.Example = `  aihunter-x scan example.com`
		}
		cmd.Parent().HelpFunc()(cmd, a)
	})
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

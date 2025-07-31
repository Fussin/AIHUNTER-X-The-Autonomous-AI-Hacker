package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/user/aihunter-x/aihunter-core-cli/cmd/aihunter-x/recon"
	"github.com/user/aihunter-x/aihunter-core-cli/cmd/aihunter-x/scan"
	"github.com/user/aihunter-x/aihunter-core-cli/cmd/aihunter-x/validate"
	"github.com/user/aihunter-x/aihunter-core-cli/core/dispatcher"
	"github.com/user/aihunter-x/aihunter-core-cli/utils"
)

var (
	rootCmd = &cobra.Command{
		Use:   "aihunter-x",
		Short: "AIHUNTER-X is an autonomous AI hacker.",
		Long:  `A fully autonomous AI hacker that tests lakhs of company domains/web apps per run.`,
		TraverseChildren: true,
		SilenceErrors:    true,
		SilenceUsage:     true,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aihunter-x.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.SuggestFor = []string{"scan", "recon", "validate", "completion", "help"}

	dispatcher.Register(&scan.Plugin{})
	dispatcher.Register(&recon.Plugin{})
	dispatcher.Register(&validate.Plugin{})

	dispatcher.AddCommands(rootCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:

Bash:

  $ source <(aihunter-x completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ aihunter-x completion bash > /etc/bash_completion.d/aihunter-x
  # macOS:
  $ aihunter-x completion bash > /usr/local/etc/bash_completion.d/aihunter-x

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ aihunter-x completion zsh > "${fpath[1]}/_aihunter-x"

  # You will need to start a new shell for this setup to take effect.

Fish:

  $ aihunter-x completion fish | source

  # To load completions for each session, execute once:
  $ aihunter-x completion fish > ~/.config/fish/completions/aihunter-x.fish

PowerShell:

  PS> aihunter-x completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> aihunter-x completion powershell > aihunter-x.ps1
  # and source this file from your PowerShell profile.
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			}
		},
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aihunter-x.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	dispatcher.Register(&scan.Plugin{})
	dispatcher.Register(&recon.Plugin{})
	dispatcher.Register(&validate.Plugin{})

	dispatcher.AddCommands(rootCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:

Bash:

  $ source <(aihunter-x completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ aihunter-x completion bash > /etc/bash_completion.d/aihunter-x
  # macOS:
  $ aihunter-x completion bash > /usr/local/etc/bash_completion.d/aihunter-x

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ aihunter-x completion zsh > "${fpath[1]}/_aihunter-x"

  # You will need to start a new shell for this setup to take effect.

Fish:

  $ aihunter-x completion fish | source

  # To load completions for each session, execute once:
  $ aihunter-x completion fish > ~/.config/fish/completions/aihunter-x.fish

PowerShell:

  PS> aihunter-x completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> aihunter-x completion powershell > aihunter-x.ps1
  # and source this file from your PowerShell profile.
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			}
		},
	})
}

func main() {
	Execute()
}

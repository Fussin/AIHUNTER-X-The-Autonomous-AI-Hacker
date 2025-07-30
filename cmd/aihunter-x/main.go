package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/user/aihunter-x/cmd/aihunter-x/recon"
	"github.com/user/aihunter-x/cmd/aihunter-x/scan"
	"github.com/user/aihunter-x/cmd/aihunter-x/validate"
	"github.com/user/aihunter-x/internal/config"
	"github.com/user/aihunter-x/internal/logger"
	"github.com/user/aihunter-x/internal/plugin"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "aihunter-x",
		Short: "AIHUNTER-X is an autonomous AI hacker.",
		Long:  `A fully autonomous AI hacker that tests lakhs of company domains/web apps per run.`,
		PersistentPreRun: initConfig,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aihunter-x.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	plugin.Register(scan.NewScanCmd)
	plugin.Register(recon.NewReconCmd)
	plugin.Register(validate.NewValidateCmd)

	plugin.AddCommands(rootCmd)
}

func initConfig(cmd *cobra.Command, args []string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName(".aihunter-x")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	verbose, _ := cmd.Flags().GetBool("verbose")
	logger.New(cfg.Log, verbose)
}

func main() {
	Execute()
}

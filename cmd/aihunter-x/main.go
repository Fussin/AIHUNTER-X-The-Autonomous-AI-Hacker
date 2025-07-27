package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/aihunter-x/internal/config"
	"github.com/user/aihunter-x/internal/engine"
	"github.com/user/aihunter-x/internal/logger"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "aihunter-x",
		Short: "AIHUNTER-X is an autonomous AI hacker.",
		Long:  `A fully autonomous AI hacker that tests lakhs of company domains/web apps per run.`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Load(cfgFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			logger.New(cfg.Log)

			engine, err := engine.NewEngine(cfg)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if err := engine.Start(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aihunter-x.yaml)")
}

func initConfig() {
	// Don't forget to read in the config file here!
}

func main() {
	Execute()
}

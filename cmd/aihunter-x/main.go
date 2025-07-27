package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/aihunter-x/internal/engine"
)

var rootCmd = &cobra.Command{
	Use:   "aihunter-x",
	Short: "AIHUNTER-X is an autonomous AI hacker.",
	Long:  `A fully autonomous AI hacker that tests lakhs of company domains/web apps per run.`,
	Run: func(cmd *cobra.Command, args []string) {
		engine, err := engine.NewEngine()
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

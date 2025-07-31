package core

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	configloader "github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
	"github.com/user/aihunter-x/aihunter-core-cli/logger/logger"
)

func InitConfig(cmd *cobra.Command, args []string) error {
	if cmd.Name() == "completion" {
		return nil
	}
	cfgFile, _ := cmd.Flags().GetString("config")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName(".aihunter-x")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	var cfg configloader.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	dump, _ := cmd.Flags().GetBool("dump")
	if !dump {
		if err := Validate(&cfg); err != nil {
			return err
		}
	}

	verbose, _ := cmd.Flags().GetBool("verbose")
	logger.New(cfg.Log, verbose)
	return nil
}

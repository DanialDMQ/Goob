package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	Protocol      string
	LastNChars    string
}

const cfgFile string = "minerConfig.json"

func LoadConfig(cmd *cobra.Command) (*Config, error) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".2_go_layout" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".2_go_layout")
	}
	// Defaults
	viper.SetDefault("Protocol", "tcp")
	viper.SetDefault("ServerAddress", "127.0.0.1:8081")
	viper.SetDefault("LastNChars", "000000")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	var config Config
	err := viper.Unmarshal(&config)

	return &config, err
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "muxi <layout>",
	Short: "Tmux layout automation",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Create muxi default layouts directory
	home := os.Getenv("HOME")
	muxiLayoutPath := fmt.Sprintf("%s/.muxi", home)
	if _, err := os.Stat(muxiLayoutPath); os.IsNotExist(err) {
		os.Mkdir(muxiLayoutPath, 0777)
	}

	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".muxi")                    // name of config file (without extension)
	viper.AddConfigPath("$HOME")                    // adding home directory as first search path
	viper.SetDefault("layoutsPath", muxiLayoutPath) // default muxi layout directory
	viper.AutomaticEnv()                            // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

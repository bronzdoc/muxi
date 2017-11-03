package cmd

import (
	"fmt"
	"github.com/bronzdoc/muxi/layout"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var show string
var edit string
var list bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "muxi <layout>",
	Short: "Tmux layout automation",

	Run: func(cmd *cobra.Command, args []string) {

		if list {
			for _, layoutName := range layout.List() {
				fmt.Printf("⚫ %s\n", layoutName)
			}
		}

		if show != "" {
			fmt.Println(string(showLayout(show)))
		}

		if edit != "" {
			if err := layout.Edit(edit); err != nil {
				fmt.Println(err)
			}
		}

		if len(args) > 0 {
			createLayout(args[0])
		}
	},
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
	RootCmd.Flags().StringVarP(&show, "show", "s", "", "show the content of a layout")
	RootCmd.Flags().StringVarP(&edit, "edit", "e", "", "edit the content of a layout")
	RootCmd.Flags().BoolVarP(&list, "list", "l", false, "list layouts")
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

func createLayout(layoutName string) {
	muxi_layout := layout.New(layoutName)

	if err := muxi_layout.Parse(); err != nil {
		fmt.Println(err)
	}

	muxi_layout.Create()
}

func showLayout(layoutName string) []byte {
	muxi_layout := layout.New(layoutName)

	if err := muxi_layout.Parse(); err != nil {
		fmt.Println(err)
	}

	return muxi_layout.RawContent()
}

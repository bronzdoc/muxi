package cmd

import (
	"fmt"
	"os"

	"github.com/bronzdoc/muxi/layout"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a tmux session using a muxi layout",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("muxi: no layout to edit given, see muxi edit --help")
			os.Exit(1)
		}

		layoutName := args[0]

		muxiLayout := layout.New(layoutName)

		if err := muxiLayout.Parse(); err != nil {
			fmt.Printf("muxi: %s", err)
			os.Exit(1)
		}

		muxiLayout.Create()
	},
}

func init() {
	RootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

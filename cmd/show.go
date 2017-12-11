package cmd

import (
	"fmt"
	"os"

	"github.com/bronzdoc/muxi/layout"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show <layout>",
	Short: "Show the content of a muxi layout",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("no layout to edit given, see muxi edit --help")
			os.Exit(1)
		}

		layoutName := args[0]

		muxiLayout := layout.New(layoutName)

		layoutContent, err := muxiLayout.RawContent()
		if err != nil {
			fmt.Println(errors.Wrap(err, "Error getting layout content"))
			os.Exit(1)
		}

		fmt.Println(string(layoutContent))
	},
}

func init() {
	RootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

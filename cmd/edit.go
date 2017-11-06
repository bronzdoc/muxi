package cmd

import (
	"fmt"
	"os"

	"github.com/bronzdoc/muxi/layout"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit <layout>",
	Short: "Edit a layout content",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("muxi: no layout to edit given, see muxi edit --help")
			os.Exit(1)
		}

		layoutName := args[0]

		if err := layout.Edit(layoutName); err != nil {
			fmt.Printf("muxi: %s\n", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

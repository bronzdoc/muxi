package cmd

import (
	"fmt"

	"github.com/bronzdoc/muxi/layout"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all muxi layouts",
	Run: func(cmd *cobra.Command, args []string) {
		if len(layout.List()) <= 0 {
			fmt.Println("no layouts found")
			return
		}

		for _, layoutName := range layout.List() {
			fmt.Printf("⚫ %s\n", layoutName)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

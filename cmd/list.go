package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Lists Passbolt Entitys",
	Long:    `Lists Passbolt Entitys`,
	Aliases: []string{"index", "ls", "filter", "search"},
}

func init() {
	rootCmd.AddCommand(listCmd)

}

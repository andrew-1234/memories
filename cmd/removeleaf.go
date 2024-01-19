package cmd

import (
	"github.com/andrew-1234/memories/internal/leaf"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var removeleafCmd = &cobra.Command{
	Use:   "removeleaf",
	Short: "remove a leaf; the file is deleted",
	Run: func(cmd *cobra.Command, args []string) {
		leaf.RemoveLeaf(JournalPath)
	},
}

func init() {
	rootCmd.AddCommand(removeleafCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	"os"

	"github.com/andrew-1234/memories/internal/leaf"
	"github.com/spf13/cobra"
)

var removenoteCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a note; the file is deleted",
	RunE: func(cmd *cobra.Command, args []string) error {
		return (leaf.RemoveNote(JournalPath, os.Stdin, os.Stdout))
	},
}

func init() {
	rootCmd.AddCommand(removenoteCmd)
}

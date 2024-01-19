package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/andrew-1234/memories/internal/leaf"
	"github.com/spf13/cobra"
)

var newleafCmd = &cobra.Command{
	Use:   "newleaf",
	Short: "Add a new memory entry",
	Long:  `This is not a long description.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Enter your journal entry:")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		// Call AddEntry from the journal package
		err := leaf.NewLeaf(text)
		if err != nil {
			fmt.Println("Error adding entry:", err)
			return
		}
		fmt.Println("Entry added successfully")
	},
}

func init() {
	rootCmd.AddCommand(newleafCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addentryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addentryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

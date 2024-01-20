package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/andrew-1234/memories/internal/leaf"
	"github.com/spf13/cobra"
)

var createnoteCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new note (memory entry)",
	Long:  `This is not a long description.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		input, err := readBuffer()
		if err != nil {
			fmt.Println("Error reading input:", err)
			return err
		}

		err = leaf.CreateNote(input, JournalPath)
		if err != nil {
			fmt.Println("Error adding entry:", err)
			return err
		}

		fmt.Println("Note created successfully")
		return err
	},
}

func init() {
	rootCmd.AddCommand(createnoteCmd)
}

func readBuffer() (string, error) {
	fmt.Println("Enter text:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return input, err
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "testing",
	Run: func(cmd *cobra.Command, args []string) {
		playground()
	},
}

func playground() {
	fmt.Println("Hello, testing")
}

func init() {
	rootCmd.AddCommand(playCmd)
}

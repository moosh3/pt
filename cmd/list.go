package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Here is a list of tasks or pomos")
	},
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var describe = &cobra.Command{
	Use:   "describe",
	Short: "Describe a pomotodo object",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Description goes here")
	},
}

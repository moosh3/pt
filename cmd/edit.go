package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var edit = &cobra.Command{
	Use:   "edit",
	Short: "Edit the metadata tied to a task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Opening pomo metadata for editting")
	},
}

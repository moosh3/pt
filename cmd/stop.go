package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stop = &cobra.Command{
	Use:   "stop",
	Short: "Stop an active pomodoro session",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Stopping")
	},
}

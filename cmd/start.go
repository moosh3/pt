package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var start = &cobra.Command{
	Use:   "start",
	Short: "Start a pomodoro session with a task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting pomo")
	},
}

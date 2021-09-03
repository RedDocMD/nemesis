package cmd

import (
	"fmt"

	"github.com/RedDocMD/nemesis/event"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all event reminders",
	Run: func(cmd *cobra.Command, args []string) {
		event.SortByDate(events)
		for _, ev := range events {
			fmt.Println(ev)
		}
	},
}

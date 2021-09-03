package cmd

import (
	"fmt"

	"github.com/RedDocMD/nemesis/event"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all event reminders",
	Run: func(cmd *cobra.Command, args []string) {
		event.SortByDate(events)
		if len(events) == 0 {
			color.Red("No events!")
		} else {
			for _, ev := range events {
				fmt.Println(ev)
			}
		}
	},
}

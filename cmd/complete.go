package cmd

import (
	"time"

	"github.com/RedDocMD/nemesis/event"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Remove events whose times are over",
	Long: `This command checks events against the current time
and removes the ones whose completion times are before.`,
	Run: func(cmd *cobra.Command, args []string) {
		event.SortByDate(events)
		currentTime := time.Now()
		threshold := 0
		for ; threshold < len(events); threshold++ {
			eventTime := events[threshold].When
			if eventTime.Equal(currentTime) || eventTime.After(currentTime) {
				break
			}
		}
		events = events[threshold:]
	},
}

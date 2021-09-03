package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/RedDocMD/nemesis/event"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var yellow = color.New(color.FgYellow).Add(color.Bold).SprintFunc()

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an event",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(events) == 0 {
			color.Red("No events!")
			return nil
		}

		event.SortByDate(events)
		eventNames := make([]string, len(events))
		for i, ev := range events {
			eventNames[i] = ev.Name
		}

		eventIdx := 0
		eventChoice := &survey.Select{
			Message: "Choose an event",
			Options: eventNames,
		}
		err := survey.AskOne(eventChoice, &eventIdx)
		if err != nil {
			return errors.WithMessage(err, "failed to delete event")
		}
		name := events[eventIdx].Name
		events = append(events[:eventIdx], events[eventIdx+1:]...)
		fmt.Println("Deleted", yellow(name))
		return nil
	},
}

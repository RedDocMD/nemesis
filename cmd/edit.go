package cmd

import (
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/RedDocMD/nemesis/event"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type changeKind = int

const (
	changeName changeKind = iota
	changeType
	changeDate
	changeTime
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an event reminder",
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
			return errors.WithMessage(err, "failed to edit event")
		}
		fmt.Println("Selected event:", events[eventIdx])

		var changeIndices []int
		changeChoice := &survey.MultiSelect{
			Message: "Choose all the parameters to change:",
			Options: []string{"Name", "Type", "Date", "Time"},
		}
		err = survey.AskOne(changeChoice, &changeIndices)
		if err != nil {
			return errors.WithMessage(err, "failed to edit event")
		}

		for _, idx := range changeIndices {
			changeEventParam(&events[eventIdx], idx)
		}
		return nil
	},
}

func changeEventParam(event *event.Event, param changeKind) error {
	var err error
	switch param {
	case changeName:
		name := ""
		err = survey.AskOne(createQuestions[param].Prompt, &name)
		if err != nil {
			return errors.WithMessage(err, "failed to get name")
		}
		event.Name = name
	case changeType:
		kind := 0
		err = survey.AskOne(createQuestions[param].Prompt, &kind)
		if err != nil {
			return errors.WithMessage(err, "failed to get kind")
		}
		event.Kind = kind
	case changeDate:
		dateStr := ""
		err = survey.AskOne(createQuestions[param].Prompt, &dateStr)
		if err != nil {
			return errors.WithMessage(err, "failed to get date")
		}
		layout := "2 Jan, 06"
		partDate, err := time.Parse(layout, dateStr)
		if err != nil {
			return errors.WithMessage(err, "failed to parse date")
		}
		oldDate := event.When
		newDate := time.Date(partDate.Hour(),
			partDate.Month(),
			partDate.Day(),
			oldDate.Hour(),
			oldDate.Minute(),
			oldDate.Second(),
			oldDate.Nanosecond(),
			oldDate.Location())
		event.When = newDate
	case changeTime:
		timeStr := ""
		err = survey.AskOne(createQuestions[param].Prompt, &timeStr)
		if err != nil {
			return errors.WithMessage(err, "failed to get date")
		}
		layout := "15:04"
		partDate, err := time.Parse(layout, timeStr)
		if err != nil {
			return errors.WithMessage(err, "failed to parse date")
		}
		oldDate := event.When
		newDate := time.Date(partDate.Hour(),
			oldDate.Month(),
			oldDate.Day(),
			partDate.Hour(),
			partDate.Minute(),
			oldDate.Second(),
			oldDate.Nanosecond(),
			oldDate.Location())
		event.When = newDate
	}
	return nil
}

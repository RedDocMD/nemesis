package cmd

import (
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/RedDocMD/nemesis/event"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new event reminder",
	Long: `Create a new event reminder.
Events may be duplicated and will not be guarded against this.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		event, err := createEvent()
		if err != nil {
			return err
		}
		events = append(events, *event)
		fmt.Println("Created", *event)
		return nil
	},
}

var qs = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "What is the name of the event?"},
		Validate: survey.Required,
	},
	{
		Name: "kind",
		Prompt: &survey.Select{
			Message: "Choose the event kind:",
			Options: event.EventKindStrings(),
		},
	},
	{
		Name: "date",
		Prompt: &survey.Input{
			Message: "Enter event date (Date Month, Year)",
		},
		Validate: func(ans interface{}) error {
			dateStr, ok := ans.(string)
			if !ok {
				return errors.New("survey don't work no more")
			}
			layout := "2 Jan, 06"
			_, err := time.Parse(layout, dateStr)
			if err != nil {
				return errors.New("not a valid date")
			}
			return nil
		},
	},
	{
		Name: "time",
		Prompt: &survey.Input{
			Message: "Enter event time (HH:MM)",
		},
		Validate: func(ans interface{}) error {
			timeStr, ok := ans.(string)
			if !ok {
				return errors.New("survey don't work no more")
			}
			layout := "15:04"
			_, err := time.Parse(layout, timeStr)
			if err != nil {
				return errors.New("not a valid time")
			}
			return nil
		},
	},
}

func createEvent() (*event.Event, error) {
	answers := struct {
		Name string
		Kind event.EventKind
		Date string
		Time string
	}{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create event")
	}
	layout := "15:04 2 Jan, 06 -0700"
	dateTimeStr := fmt.Sprintf("%s %s +0530", answers.Time, answers.Date)
	when, _ := time.Parse(layout, dateTimeStr)
	event := &event.Event{
		Name: answers.Name,
		Kind: answers.Kind,
		When: when,
	}
	return event, nil
}

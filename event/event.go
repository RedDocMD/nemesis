package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

type EventKind = int

const (
	AssignmentEvent EventKind = iota
	ExamEvent
)

type Event struct {
	Name string
	Kind EventKind
	When time.Time
}

func GetEvents(path string) ([]Event, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil
	}
	defer file.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get events")
	}
	var events []Event
	err = json.Unmarshal(buf.Bytes(), &events)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to parse events")
	}
	return events, nil
}

func DumpEvents(path string, events []Event) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.WithMessage(err, "failed to dump events")
	}
	defer file.Close()
	content, err := json.Marshal(events)
	if err != nil {
		return errors.WithMessage(err, "failed to marshall events")
	}
	_, err = file.Write(content)
	if err != nil {
		return errors.WithMessage(err, "failed to dump events")
	}
	return nil
}

func EventKindStrings() []string {
	return []string{"Assignment", "Exam"}
}

func (ev Event) KindName() string {
	return EventKindStrings()[ev.Kind]
}

func (ev Event) String() string {
	yellow := color.New(color.FgYellow).Add(color.Bold).SprintFunc()
	red := color.New(color.FgRed).Add(color.Bold).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	when := ev.When
	date := fmt.Sprintf("%d %s, %d", when.Day(), when.Month(), when.Year())
	time := fmt.Sprintf("%d:%d", when.Hour(), when.Minute())

	return fmt.Sprintf("%s %s on %s, %s", yellow(ev.Name), ev.KindName(), cyan(time), red(date))
}

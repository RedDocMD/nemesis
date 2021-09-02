package event

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/pkg/errors"
)

type EventKind = int

const (
	TestEvent EventKind = iota
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
		return nil, errors.WithMessage(err, "failed to get events")
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

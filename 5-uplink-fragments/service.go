package flow30

import (
	"errors"
	"telemetry-signal-routing-service/internal/state30"
)

var ErrBadItem = errors.New("invalid fragment")

func Process(tracker *state30.Tracker, frames []string) (err error) {
	for _, frame := range frames {
		resource, openErr := tracker.Open()
		if openErr != nil {
			return openErr
		}
		defer func() { err = resource.Finish(err) }()
		if frame == "bad" {
			return ErrBadItem
		}
	}
	return nil
}

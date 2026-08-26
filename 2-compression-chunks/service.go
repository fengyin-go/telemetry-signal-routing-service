package flow27

import (
	"errors"
	"telemetry-signal-routing-service/internal/state27"
)

var ErrBadItem = errors.New("invalid chunk")

func Process(tracker *state27.Tracker, frames []string) (err error) {
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

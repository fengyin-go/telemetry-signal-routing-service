package flow29

import (
	"errors"
	"telemetry-signal-routing-service/internal/state29"
)

var ErrBadItem = errors.New("invalid block")

func Process(tracker *state29.Tracker, frames []string) (err error) {
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

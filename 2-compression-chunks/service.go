package flow27

import (
	"errors"
	"telemetry-signal-routing-service/internal/state27"
)

var ErrBadItem = errors.New("invalid chunk")

func Process(tracker *state27.Tracker, frames []string) error {
	for _, frame := range frames {
		err := func() (err error) {
			resource, openErr := tracker.Open()
			if openErr != nil {
				return openErr
			}
			defer func() { err = resource.Finish(err) }()
			if frame == "bad" {
				return ErrBadItem
			}
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

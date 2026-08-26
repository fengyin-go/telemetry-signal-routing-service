package flow25_test

import (
	"errors"
	"telemetry-signal-routing-service/internal/state25"
	flow25 "telemetry-signal-routing-service/y-calibration-frames"
	"testing"
)

func TestY(t *testing.T) {
	frames := []string{"ok", "bad"}
	tracker := state25.NewTracker(1)
	err := flow25.Process(tracker, frames)
	if !errors.Is(err, flow25.ErrBadItem) || tracker.OpenCount() != 0 {
		t.Fatalf("calibration-frames batch lost its item error or retained an open resource")
	}
}

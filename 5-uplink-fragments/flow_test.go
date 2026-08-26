package flow30_test

import (
	"errors"
	flow30 "telemetry-signal-routing-service/5-uplink-fragments"
	"telemetry-signal-routing-service/internal/state30"
	"testing"
)

func Test5(t *testing.T) {
	frames := []string{"ok", "bad"}
	tracker := state30.NewTracker(1)
	err := flow30.Process(tracker, frames)
	if !errors.Is(err, flow30.ErrBadItem) || tracker.OpenCount() != 0 {
		t.Fatalf("uplink-fragments batch lost its item error or retained an open resource")
	}
}

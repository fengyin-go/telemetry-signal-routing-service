package flow28_test

import (
	"errors"
	flow28 "telemetry-signal-routing-service/3-relay-segments"
	"telemetry-signal-routing-service/internal/state28"
	"testing"
)

func Test3(t *testing.T) {
	frames := []string{"ok", "bad"}
	tracker := state28.NewTracker(1)
	err := flow28.Process(tracker, frames)
	if !errors.Is(err, flow28.ErrBadItem) || tracker.OpenCount() != 0 {
		t.Fatalf("relay-segments batch lost its item error or retained an open resource")
	}
}

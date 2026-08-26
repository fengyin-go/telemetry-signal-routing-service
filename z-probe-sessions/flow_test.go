package flow26_test

import (
	"errors"
	"telemetry-signal-routing-service/internal/state26"
	flow26 "telemetry-signal-routing-service/z-probe-sessions"
	"testing"
)

func TestZ(t *testing.T) {
	frames := []string{"ok", "bad"}
	tracker := state26.NewTracker(1)
	err := flow26.Process(tracker, frames)
	if !errors.Is(err, flow26.ErrBadItem) || tracker.OpenCount() != 0 {
		t.Fatalf("probe-sessions batch lost its item error or retained an open resource")
	}
}

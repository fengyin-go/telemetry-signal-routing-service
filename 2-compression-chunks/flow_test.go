package flow27_test

import (
	"errors"
	flow27 "telemetry-signal-routing-service/2-compression-chunks"
	"telemetry-signal-routing-service/internal/state27"
	"testing"
)

func Test2(t *testing.T) {
	frames := []string{"ok", "bad"}
	tracker := state27.NewTracker(1)
	err := flow27.Process(tracker, frames)
	if !errors.Is(err, flow27.ErrBadItem) || tracker.OpenCount() != 0 {
		t.Fatalf("compression-chunks batch lost its item error or retained an open resource")
	}
}

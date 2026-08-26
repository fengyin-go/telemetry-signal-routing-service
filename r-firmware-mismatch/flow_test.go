package flow18_test

import (
	"errors"
	"telemetry-signal-routing-service/internal/state18"
	flow18 "telemetry-signal-routing-service/r-firmware-mismatch"
	"testing"
)

func TestR(t *testing.T) {
	source := state18.NewSource(&state18.Rejected{Reason: "firmware"}, nil)
	err := flow18.Forward(source, 2)
	var rejected *state18.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent firmware-mismatch response was retried or lost its typed rejection")
	}
}

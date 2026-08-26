package flow17_test

import (
	"errors"
	"telemetry-signal-routing-service/internal/state17"
	flow17 "telemetry-signal-routing-service/q-checksum-refusal"
	"testing"
)

func TestQ(t *testing.T) {
	source := state17.NewSource(&state17.Rejected{Reason: "checksum"}, nil)
	err := flow17.Forward(source, 2)
	var rejected *state17.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent checksum-refusal response was retried or lost its typed rejection")
	}
}

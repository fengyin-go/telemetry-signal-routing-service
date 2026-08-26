package flow13_test

import (
	"errors"
	"telemetry-signal-routing-service/internal/state13"
	flow13 "telemetry-signal-routing-service/m-decoder-rejection"
	"testing"
)

func TestM(t *testing.T) {
	source := state13.NewSource(&state13.Rejected{Reason: "bad-codec"}, nil)
	err := flow13.Forward(source, 2)
	var rejected *state13.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent decoder-rejection response was retried or lost its typed rejection")
	}
}

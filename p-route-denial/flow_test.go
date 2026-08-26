package flow16_test

import (
	"errors"
	"telemetry-signal-routing-service/internal/state16"
	flow16 "telemetry-signal-routing-service/p-route-denial"
	"testing"
)

func TestP(t *testing.T) {
	source := state16.NewSource(&state16.Rejected{Reason: "route"}, nil)
	err := flow16.Forward(source, 2)
	var rejected *state16.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent route-denial response was retried or lost its typed rejection")
	}
}

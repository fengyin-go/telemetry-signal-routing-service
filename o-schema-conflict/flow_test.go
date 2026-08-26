package flow15_test

import (
	"errors"
	"telemetry-signal-routing-service/internal/state15"
	flow15 "telemetry-signal-routing-service/o-schema-conflict"
	"testing"
)

func TestO(t *testing.T) {
	source := state15.NewSource(&state15.Rejected{Reason: "schema"}, nil)
	err := flow15.Forward(source, 2)
	var rejected *state15.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent schema-conflict response was retried or lost its typed rejection")
	}
}

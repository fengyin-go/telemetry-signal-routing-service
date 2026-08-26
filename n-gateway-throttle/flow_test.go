package flow14_test

import (
	"errors"
	"telemetry-signal-routing-service/internal/state14"
	flow14 "telemetry-signal-routing-service/n-gateway-throttle"
	"testing"
)

func TestN(t *testing.T) {
	source := state14.NewSource(&state14.Rejected{Reason: "quota"}, nil)
	err := flow14.Forward(source, 2)
	var rejected *state14.Rejected
	if source.Calls() != 1 || !errors.As(err, &rejected) {
		t.Fatalf("permanent gateway-throttle response was retried or lost its typed rejection")
	}
}

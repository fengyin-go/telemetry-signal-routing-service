package flow19_test

import (
	"telemetry-signal-routing-service/internal/state19"
	flow19 "telemetry-signal-routing-service/s-threshold-gate"
	"testing"
)

func callDisabledGate(gate state19.Gate, payload string) (err error, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	err = flow19.Accept(gate, payload)
	return err, false
}

func TestS(t *testing.T) {
	gate := state19.NewGate(false)
	err, panicked := callDisabledGate(gate, "threshold")
	if panicked || err != nil {
		t.Fatalf("disabled threshold-gate gate rejected or panicked on an accepted signal")
	}
}

package flow20_test

import (
	"telemetry-signal-routing-service/internal/state20"
	flow20 "telemetry-signal-routing-service/t-quality-screen"
	"testing"
)

func callDisabledGate(gate state20.Gate, payload string) (err error, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	err = flow20.Accept(gate, payload)
	return err, false
}

func TestT(t *testing.T) {
	gate := state20.NewGate(false)
	err, panicked := callDisabledGate(gate, "quality")
	if panicked || err != nil {
		t.Fatalf("disabled quality-screen gate rejected or panicked on an accepted signal")
	}
}

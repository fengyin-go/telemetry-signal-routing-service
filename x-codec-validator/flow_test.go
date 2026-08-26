package flow24_test

import (
	"telemetry-signal-routing-service/internal/state24"
	flow24 "telemetry-signal-routing-service/x-codec-validator"
	"testing"
)

func callDisabledGate(gate state24.Gate, payload string) (err error, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	err = flow24.Accept(gate, payload)
	return err, false
}

func TestX(t *testing.T) {
	gate := state24.NewGate(false)
	err, panicked := callDisabledGate(gate, "codec")
	if panicked || err != nil {
		t.Fatalf("disabled codec-validator gate rejected or panicked on an accepted signal")
	}
}

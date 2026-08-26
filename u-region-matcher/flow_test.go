package flow21_test

import (
	"telemetry-signal-routing-service/internal/state21"
	flow21 "telemetry-signal-routing-service/u-region-matcher"
	"testing"
)

func callDisabledGate(gate state21.Gate, payload string) (err error, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	err = flow21.Accept(gate, payload)
	return err, false
}

func TestU(t *testing.T) {
	gate := state21.NewGate(false)
	err, panicked := callDisabledGate(gate, "region")
	if panicked || err != nil {
		t.Fatalf("disabled region-matcher gate rejected or panicked on an accepted signal")
	}
}

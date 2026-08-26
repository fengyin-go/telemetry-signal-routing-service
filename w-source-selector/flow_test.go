package flow23_test

import (
	"telemetry-signal-routing-service/internal/state23"
	flow23 "telemetry-signal-routing-service/w-source-selector"
	"testing"
)

func callDisabledGate(gate state23.Gate, payload string) (err error, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	err = flow23.Accept(gate, payload)
	return err, false
}

func TestW(t *testing.T) {
	gate := state23.NewGate(false)
	err, panicked := callDisabledGate(gate, "source")
	if panicked || err != nil {
		t.Fatalf("disabled source-selector gate rejected or panicked on an accepted signal")
	}
}

package flow22_test

import (
	"telemetry-signal-routing-service/internal/state22"
	flow22 "telemetry-signal-routing-service/v-signal-filter"
	"testing"
)

func callDisabledGate(gate state22.Gate, payload string) (err error, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	err = flow22.Accept(gate, payload)
	return err, false
}

func TestV(t *testing.T) {
	gate := state22.NewGate(false)
	err, panicked := callDisabledGate(gate, "filter")
	if panicked || err != nil {
		t.Fatalf("disabled signal-filter gate rejected or panicked on an accepted signal")
	}
}

package flow07_test

import (
	flow07 "telemetry-signal-routing-service/g-spectral-bands"
	"telemetry-signal-routing-service/internal/state07"
	"testing"
)

func TestG(t *testing.T) {
	input := []string{"alpha"}
	store := &state07.Store{}
	flow07.Capture(store, input)
	input[0] = "later-input"
	first := flow07.Read(store)
	first[0] = "later-read"
	second := flow07.Read(store)
	if len(second) != 1 || second[0] != "alpha" {
		t.Fatalf("captured spectral-bands values changed after later ownership mutations")
	}
}

package flow10_test

import (
	"telemetry-signal-routing-service/internal/state10"
	flow10 "telemetry-signal-routing-service/j-packet-labels"
	"testing"
)

func TestJ(t *testing.T) {
	input := []string{"stable"}
	store := &state10.Store{}
	flow10.Capture(store, input)
	input[0] = "later-input"
	first := flow10.Read(store)
	first[0] = "later-read"
	second := flow10.Read(store)
	if len(second) != 1 || second[0] != "stable" {
		t.Fatalf("captured packet-labels values changed after later ownership mutations")
	}
}

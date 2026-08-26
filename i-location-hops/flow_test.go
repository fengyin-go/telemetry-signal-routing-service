package flow09_test

import (
	flow09 "telemetry-signal-routing-service/i-location-hops"
	"telemetry-signal-routing-service/internal/state09"
	"testing"
)

func TestI(t *testing.T) {
	input := []string{"sector-a"}
	store := &state09.Store{}
	flow09.Capture(store, input)
	input[0] = "later-input"
	first := flow09.Read(store)
	first[0] = "later-read"
	second := flow09.Read(store)
	if len(second) != 1 || second[0] != "sector-a" {
		t.Fatalf("captured location-hops values changed after later ownership mutations")
	}
}

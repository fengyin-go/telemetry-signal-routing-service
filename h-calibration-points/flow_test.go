package flow08_test

import (
	flow08 "telemetry-signal-routing-service/h-calibration-points"
	"telemetry-signal-routing-service/internal/state08"
	"testing"
)

func TestH(t *testing.T) {
	input := []string{"north"}
	store := &state08.Store{}
	flow08.Capture(store, input)
	input[0] = "later-input"
	first := flow08.Read(store)
	first[0] = "later-read"
	second := flow08.Read(store)
	if len(second) != 1 || second[0] != "north" {
		t.Fatalf("captured calibration-points values changed after later ownership mutations")
	}
}

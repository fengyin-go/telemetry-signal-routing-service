package flow11_test

import (
	"telemetry-signal-routing-service/internal/state11"
	flow11 "telemetry-signal-routing-service/k-anomaly-tags"
	"testing"
)

func TestK(t *testing.T) {
	input := []string{"baseline"}
	store := &state11.Store{}
	flow11.Capture(store, input)
	input[0] = "later-input"
	first := flow11.Read(store)
	first[0] = "later-read"
	second := flow11.Read(store)
	if len(second) != 1 || second[0] != "baseline" {
		t.Fatalf("captured anomaly-tags values changed after later ownership mutations")
	}
}

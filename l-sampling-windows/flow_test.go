package flow12_test

import (
	"telemetry-signal-routing-service/internal/state12"
	flow12 "telemetry-signal-routing-service/l-sampling-windows"
	"testing"
)

func TestL(t *testing.T) {
	input := []string{"window-a"}
	store := &state12.Store{}
	flow12.Capture(store, input)
	input[0] = "later-input"
	first := flow12.Read(store)
	first[0] = "later-read"
	second := flow12.Read(store)
	if len(second) != 1 || second[0] != "window-a" {
		t.Fatalf("captured sampling-windows values changed after later ownership mutations")
	}
}

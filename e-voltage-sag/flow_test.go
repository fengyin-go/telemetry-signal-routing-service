package flow05_test

import (
	"context"
	"errors"
	flow05 "telemetry-signal-routing-service/e-voltage-sag"
	"telemetry-signal-routing-service/internal/state05"
	"testing"
)

func TestE(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state05.Sink{}
	err := flow05.Route(ctx, sink, "voltage")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled voltage-sag signal was delivered without a cancellation result")
	}
}

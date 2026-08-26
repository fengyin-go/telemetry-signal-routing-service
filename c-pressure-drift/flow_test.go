package flow03_test

import (
	"context"
	"errors"
	flow03 "telemetry-signal-routing-service/c-pressure-drift"
	"telemetry-signal-routing-service/internal/state03"
	"testing"
)

func TestC(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state03.Sink{}
	err := flow03.Route(ctx, sink, "pressure")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled pressure-drift signal was delivered without a cancellation result")
	}
}

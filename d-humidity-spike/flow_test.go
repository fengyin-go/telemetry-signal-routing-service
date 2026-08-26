package flow04_test

import (
	"context"
	"errors"
	flow04 "telemetry-signal-routing-service/d-humidity-spike"
	"telemetry-signal-routing-service/internal/state04"
	"testing"
)

func TestD(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state04.Sink{}
	err := flow04.Route(ctx, sink, "humidity")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled humidity-spike signal was delivered without a cancellation result")
	}
}

package flow01_test

import (
	"context"
	"errors"
	flow01 "telemetry-signal-routing-service/a-thermal-burst"
	"telemetry-signal-routing-service/internal/state01"
	"testing"
)

func TestA(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state01.Sink{}
	err := flow01.Route(ctx, sink, "thermal")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled thermal-burst signal was delivered without a cancellation result")
	}
}

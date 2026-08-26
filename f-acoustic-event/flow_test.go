package flow06_test

import (
	"context"
	"errors"
	flow06 "telemetry-signal-routing-service/f-acoustic-event"
	"telemetry-signal-routing-service/internal/state06"
	"testing"
)

func TestF(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sink := &state06.Sink{}
	err := flow06.Route(ctx, sink, "acoustic")
	if !errors.Is(err, context.Canceled) || sink.Calls() != 0 {
		t.Fatalf("cancelled acoustic-event signal was delivered without a cancellation result")
	}
}

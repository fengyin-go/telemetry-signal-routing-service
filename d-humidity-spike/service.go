package flow04

import (
	"context"
	"telemetry-signal-routing-service/internal/state04"
)

func Route(ctx context.Context, sink *state04.Sink, signal string) error {
	return sink.Deliver(context.Background(), signal)
}

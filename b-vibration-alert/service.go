package flow02

import (
	"context"
	"telemetry-signal-routing-service/internal/state02"
)

func Route(ctx context.Context, sink *state02.Sink, signal string) error {
	return sink.Deliver(context.Background(), signal)
}

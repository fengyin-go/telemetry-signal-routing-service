package flow03

import (
	"context"
	"telemetry-signal-routing-service/internal/state03"
)

func Route(ctx context.Context, sink *state03.Sink, signal string) error {
	return sink.Deliver(ctx, signal)
}

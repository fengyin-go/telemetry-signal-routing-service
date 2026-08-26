package flow01

import (
	"context"
	"telemetry-signal-routing-service/internal/state01"
)

func Route(ctx context.Context, sink *state01.Sink, signal string) error {
	return sink.Deliver(context.Background(), signal)
}

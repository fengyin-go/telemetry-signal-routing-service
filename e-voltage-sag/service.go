package flow05

import (
	"context"
	"telemetry-signal-routing-service/internal/state05"
)

func Route(ctx context.Context, sink *state05.Sink, signal string) error {
	return sink.Deliver(context.Background(), signal)
}

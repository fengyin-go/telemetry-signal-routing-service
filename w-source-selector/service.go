package flow23

import "telemetry-signal-routing-service/internal/state23"

func Accept(gate state23.Gate, payload string) error {
	return gate.Validate(payload)
}

package flow23

import "telemetry-signal-routing-service/internal/state23"

func Accept(gate state23.Gate, payload string) error {
	if gate == nil {
		return nil
	}
	return gate.Validate(payload)
}

package flow24

import "telemetry-signal-routing-service/internal/state24"

func Accept(gate state24.Gate, payload string) error {
	if gate == nil {
		return nil
	}
	return gate.Validate(payload)
}

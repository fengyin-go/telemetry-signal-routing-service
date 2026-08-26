package flow20

import "telemetry-signal-routing-service/internal/state20"

func Accept(gate state20.Gate, payload string) error {
	if gate == nil {
		return nil
	}
	return gate.Validate(payload)
}

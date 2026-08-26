package flow22

import "telemetry-signal-routing-service/internal/state22"

func Accept(gate state22.Gate, payload string) error {
	if gate == nil {
		return nil
	}
	return gate.Validate(payload)
}

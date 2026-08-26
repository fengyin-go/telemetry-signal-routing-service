package flow20

import "telemetry-signal-routing-service/internal/state20"

func Accept(gate state20.Gate, payload string) error {
	return gate.Validate(payload)
}

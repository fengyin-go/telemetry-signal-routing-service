package flow19

import "telemetry-signal-routing-service/internal/state19"

func Accept(gate state19.Gate, payload string) error {
	return gate.Validate(payload)
}

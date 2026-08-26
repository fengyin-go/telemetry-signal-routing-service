package flow21

import "telemetry-signal-routing-service/internal/state21"

func Accept(gate state21.Gate, payload string) error {
	return gate.Validate(payload)
}

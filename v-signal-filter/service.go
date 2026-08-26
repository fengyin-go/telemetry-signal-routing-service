package flow22

import "telemetry-signal-routing-service/internal/state22"

func Accept(gate state22.Gate, payload string) error {
	return gate.Validate(payload)
}

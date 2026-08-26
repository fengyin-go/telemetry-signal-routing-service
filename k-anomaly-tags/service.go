package flow11

import "telemetry-signal-routing-service/internal/state11"

func Capture(store *state11.Store, values []string) {
	store.Replace(values)
}

func Read(store *state11.Store) []string { return store.Snapshot() }

package flow12

import "telemetry-signal-routing-service/internal/state12"

func Capture(store *state12.Store, values []string) {
	store.Replace(values)
}

func Read(store *state12.Store) []string { return store.Snapshot() }

package flow08

import "telemetry-signal-routing-service/internal/state08"

func Capture(store *state08.Store, values []string) {
	store.Replace(values)
}

func Read(store *state08.Store) []string { return store.Snapshot() }

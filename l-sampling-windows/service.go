package flow12

import "telemetry-signal-routing-service/internal/state12"

func Capture(store *state12.Store, values []string) {
	snapshot := append([]string(nil), values...)
	store.Replace(snapshot)
}

func Read(store *state12.Store) []string { return store.Snapshot() }

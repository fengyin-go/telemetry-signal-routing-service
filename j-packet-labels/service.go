package flow10

import "telemetry-signal-routing-service/internal/state10"

func Capture(store *state10.Store, values []string) {
	snapshot := append([]string(nil), values...)
	store.Replace(snapshot)
}

func Read(store *state10.Store) []string { return store.Snapshot() }

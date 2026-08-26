package state09

type Store struct{ values []string }

// Replace stores a private copy of values so later mutation of the caller's
// buffer (e.g. reusing it for a second round) cannot alter what was captured.
func (s *Store) Replace(values []string) {
	s.values = append([]string(nil), values...)
}

// Snapshot returns an independent copy so callers cannot mutate the store's
// internal slice through the returned value.
func (s *Store) Snapshot() []string {
	return append([]string(nil), s.values...)
}

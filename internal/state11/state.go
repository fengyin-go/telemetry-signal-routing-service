package state11

type Store struct{ values []string }

// Replace stores an independent copy of values so later mutations to the
// caller's slice (clearing, sorting, truncating) cannot alter the saved data.
func (s *Store) Replace(values []string) {
	s.values = append([]string(nil), values...)
}

// Snapshot returns a copy of the stored values so callers cannot mutate the
// store's internal state through the returned slice.
func (s *Store) Snapshot() []string {
	return append([]string(nil), s.values...)
}

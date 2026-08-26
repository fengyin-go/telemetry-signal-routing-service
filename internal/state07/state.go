package state07

type Store struct{ values []string }

// Replace takes an independent copy of values so the store is immune to the
// caller reusing or mutating the supplied input buffer afterwards.
func (s *Store) Replace(values []string) {
	s.values = append([]string(nil), values...)
}

// Snapshot returns an independent copy so mutating one read cannot leak back
// into the store or affect subsequent reads.
func (s *Store) Snapshot() []string {
	return append([]string(nil), s.values...)
}

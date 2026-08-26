package state12

type Store struct{ values []string }

// Replace stores a copy of values so later mutations to the caller's
// receive buffer cannot overwrite the saved history.
func (s *Store) Replace(values []string) {
	s.values = append([]string(nil), values...)
}

// Snapshot returns a copy of the stored history so mutations made to the
// returned slice cannot leak into a later readback.
func (s *Store) Snapshot() []string {
	return append([]string(nil), s.values...)
}

package state10

type Store struct{ values []string }

// Replace stores a copy of values, so later reuse of the caller's input
// buffer cannot mutate the cached snapshot.
func (s *Store) Replace(values []string) {
	dup := make([]string, len(values))
	copy(dup, values)
	s.values = dup
}

// Snapshot returns a copy of the stored values, so mutations the caller
// makes to the returned slice cannot alter the store's next result.
func (s *Store) Snapshot() []string {
	dup := make([]string, len(s.values))
	copy(dup, s.values)
	return dup
}

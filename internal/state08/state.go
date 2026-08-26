package state08

type Store struct{ values []string }

func (s *Store) Replace(values []string) {
	dup := make([]string, len(values))
	copy(dup, values)
	s.values = dup
}

func (s *Store) Snapshot() []string {
	dup := make([]string, len(s.values))
	copy(dup, s.values)
	return dup
}

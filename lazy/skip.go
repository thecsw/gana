package lazy

import "github.com/thecsw/gana"

type SkipT[T any] struct {
	iter Iter[T]
	n    int
}

// Make sure that SkipT implements the Iter interface.
var _ Iter[string] = (*SkipT[string])(nil)

// Skip returns an iterator that skips the first n elements of the given iterator.
func Skip[T any](iter Iter[T], n int) Iter[T] {
	return &SkipT[T]{iter, n}
}

func (s *SkipT[T]) Next() (T, bool) {
	for s.n > 0 {
		_, ok := s.iter.Next()
		if !ok {
			return gana.ZeroValue[T](), false
		}
		s.n--
	}
	return s.iter.Next()
}

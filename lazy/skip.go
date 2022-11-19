package lazy

import "github.com/thecsw/gana"

type SkipT[T any] struct {
	iter Iter[T]
	n    int
}

type SkipWhileT[T any] struct {
	iter    Iter[T]
	pred    func(T) bool
	skipped bool
}

// Make sure that SkipT and SkipWhileT implements the Iter interface.
var _ Iter[string] = (*SkipT[string])(nil)
var _ Iter[string] = (*SkipWhileT[string])(nil)

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

// SkipWhile returns an iterator that skips the first group of elements that satisfy the predicate.
func SkipWhile[T any](iter Iter[T], pred func(T) bool) Iter[T] {
	return &SkipWhileT[T]{iter, pred, false}
}

func (s *SkipWhileT[T]) Next() (T, bool) {
	for !s.skipped {
		e, ok := s.iter.Next()
		if !ok {
			return gana.ZeroValue[T](), false
		}
		if !s.pred(e) {
			s.skipped = true
			return e, true
		}
	}
	return s.iter.Next()
}

package lazy

import "github.com/thecsw/gana"

type FilterT[T any] struct {
	iter Iter[T]
	pred func(T) bool
}

// Make sure that FilterT implements the Iter interface.
var _ Iter[string] = (*FilterT[string])(nil)

// Filter returns an iterator over the elements of `iter` that satisfy the predicate `pred`.
func Filter[T any](iter Iter[T], pred func(T) bool) Iter[T] {
	return &FilterT[T]{iter, pred}
}

func (f *FilterT[T]) Next() (T, bool) {
	for {
		e, ok := f.iter.Next()
		if !ok {
			return gana.ZeroValue[T](), false
		}
		if f.pred(e) {
			return e, true
		}
	}
}

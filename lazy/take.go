package lazy

import "github.com/thecsw/gana"

type TakeT[T any] struct {
	iter      Iter[T]
	remaining int
}

// Make sure that TakeT implements the Iter interface.
var _ Iter[string] = (*TakeT[string])(nil)

// Take returns an iterator over the first `n` elements of `iter`.
func Take[T any](iter Iter[T], n int) Iter[T] {
	return &TakeT[T]{iter, n}
}

func (t *TakeT[T]) Next() (T, bool) {
	if t.remaining <= 0 {
		return gana.ZeroValue[T](), false
	}
	t.remaining--
	return t.iter.Next()
}

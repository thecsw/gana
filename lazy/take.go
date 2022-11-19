package lazy

import "github.com/thecsw/gana"

type TakeT[T any] struct {
	iter      Iter[T]
	remaining int
}

type TakeWhileT[T any] struct {
	iter Iter[T]
	pred func(T) bool
}

// Make sure that TakeT implements the Iter interface.
var _ Iter[string] = (*TakeT[string])(nil)
var _ Iter[string] = (*TakeWhileT[string])(nil)

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

// TakeWhile returns an iterator over the first group of elements that satisfy the predicate.
func TakeWhile[T any](iter Iter[T], pred func(T) bool) Iter[T] {
	return &TakeWhileT[T]{iter, pred}
}

func (t *TakeWhileT[T]) Next() (T, bool) {
	e, ok := t.iter.Next()
	if !ok {
		return gana.ZeroValue[T](), false
	}
	if !t.pred(e) {
		t.pred = func(T) bool { return false }
		return gana.ZeroValue[T](), false
	}
	return e, true
}

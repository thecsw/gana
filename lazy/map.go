package lazy

import "github.com/thecsw/gana"

type MapT[T any, U any] struct {
	iter Iter[T]
	f    func(T) U
}

// Make sure that MapT implements the Iter interface.
var _ Iter[int] = (*MapT[string, int])(nil)

// Map returns an iterator over the elements of `iter` transformed by the function `f`.
func Map[T any, U any](iter Iter[T], f func(T) U) Iter[U] {
	return &MapT[T, U]{iter, f}
}

func (m *MapT[T, U]) Next() (U, bool) {
	e, ok := m.iter.Next()
	if !ok {
		return gana.ZeroValue[U](), false
	}
	return m.f(e), true
}

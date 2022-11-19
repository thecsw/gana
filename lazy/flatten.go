package lazy

import "github.com/thecsw/gana"

type FlattenT[T any] struct {
	iter    Iter[Iter[T]]
	current Iter[T]
}

// Make sure that FlattenT implements the Iter interface.
var _ Iter[string] = (*FlattenT[string])(nil)

// FlattenIters returns an iterator over the elements of the iterators in `iter`.
func FlattenIters[T any](iter Iter[Iter[T]]) Iter[T] {
	current, ok := iter.Next()
	if !ok {
		return &EmptyT[T]{}
	}

	return &FlattenT[T]{iter, current}
}

// Flatten returns an iterator over the elements of the slices in `iter`.
func Flatten[T any](iter Iter[[]T]) Iter[T] {
	return FlattenIters(Map(iter, NewIter[T]))
}

func (f *FlattenT[T]) Next() (T, bool) {
	for {
		e, ok := f.current.Next()
		if ok {
			return e, true
		}

		f.current, ok = f.iter.Next()
		if !ok {
			return gana.ZeroValue[T](), false
		}
	}
}

package lazy

import "github.com/thecsw/gana"

// Iterator is a lazy iterator over a slice of elements.
type Iter[T any] interface {
	// Next returns the next element in the iterator.
	// If there are no more elements, it returns `gana.ZeroValue[T](), false`.
	Next() (T, bool)
}

type IterT[T any] struct {
	slice []T
	index int
}

// Make sure that the base IterT implements the Iter interface.
var _ Iter[string] = (*IterT[string])(nil)

// NewIter returns a lazy iterator over the elements of `slice`.
func NewIter[T any](slice []T) Iter[T] {
	return &IterT[T]{slice: slice, index: 0}
}

func (i *IterT[T]) Next() (T, bool) {
	if i.index >= len(i.slice) {
		return gana.ZeroValue[T](), false
	}
	i.index++
	return i.slice[i.index-1], true
}

package lazy

import (
	"github.com/thecsw/gana"
)

type ZipT[T any, U any] struct {
	left  Iter[T]
	right Iter[U]
}

// Make sure that ZipT implements the Iter interface.
var _ Iter[gana.Tuple[int, string]] = (*ZipT[int, string])(nil)

// Zip returns an iterator over the elements of `left` and `right` zipped together.
// The iterator will stop when either `left` or `right` is exhausted.
func Zip[T any, U any](left Iter[T], right Iter[U]) Iter[gana.Tuple[T, U]] {
	return &ZipT[T, U]{left, right}
}

func (z *ZipT[T, U]) Next() (gana.Tuple[T, U], bool) {
	l, ok := z.left.Next()
	if !ok {
		return gana.ZeroValue[gana.Tuple[T, U]](), false
	}
	r, ok := z.right.Next()
	if !ok {
		return gana.ZeroValue[gana.Tuple[T, U]](), false
	}
	return gana.NewTuple(l, r), true
}

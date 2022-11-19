package lazy

import "github.com/thecsw/gana"

// Fold applies a fold `f` over the elements of `iter`, starting with the initial value `init`.
func Fold[T, U any](iter Iter[T], init U, f func(U, T) U) U {
	result := init
	for {
		e, ok := iter.Next()
		if !ok {
			break
		}
		result = f(result, e)
	}
	return result
}

// Fold1 applies a fold `f` over the elements of `iter`, starting with the first element of `iter`.
// If `iter` is empty, it returns `gana.ZeroValue[T]()“.
func Fold1[T any](iter Iter[T], f func(T, T) T) T {
	e, ok := iter.Next()
	if !ok {
		return gana.ZeroValue[T]()
	}
	return Fold(iter, e, f)
}

// FoldRight applies a fold `f` over the elements of `iter`, starting with the initial value `init`.
// The elements of `iter` are traversed in reverse order.
// If `iter` is empty, it returns `init`.
func FoldRight[T, U any](iter Iter[T], init U, f func(T, U) U) U {
	result := init
	slice := Collect(iter)
	for i := len(slice) - 1; i >= 0; i-- {
		result = f(slice[i], result)
	}
	return result
}

// FoldRight1 applies a fold `f` over the elements of `iter`, starting with the last element of `iter`.
// The elements of `iter` are traversed in reverse order.
// If `iter` is empty, it returns `gana.ZeroValue[T]()`.
func FoldRight1[T any](iter Iter[T], f func(T, T) T) T {
	slice := Collect(iter)
	if len(slice) == 0 {
		return gana.ZeroValue[T]()
	}
	result := slice[len(slice)-1]
	for i := len(slice) - 2; i >= 0; i-- {
		result = f(slice[i], result)
	}
	return result
}

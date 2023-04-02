package gana

import "golang.org/x/exp/constraints"

// Min returns the minimum of two numbers.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two numbers.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Less is the default less function for any ordered type and will
// return true if a < b, false otherwise.
//
//go:inline
func Less[T constraints.Ordered](a, b T) bool { return a < b }

// DefaultLess is the default greater function for any ordered type and will
// return true if a > b, false otherwise.
//
//go:inline
func Greater[T constraints.Ordered](a, b T) bool { return a > b }

// Equals is the default equals function for any ordered type and will
// return true if a = b, false otherwise.
//
//go:inline
func Equals[T comparable](a, b T) bool { return a == b }

// Equals is the default equals function that will return a closure equals
// function to compare the value to the original wrapped val.
//
//go:inline
func EqualsClosure[T comparable](a T) func(T) bool { return func(b T) bool { return a == b } }

// Minv returns the minimum of the given values, or if no values are given, the zero value of the type.
//
//go:inline
func Minv[T constraints.Ordered](vals ...T) T {
	return Minf(Less[T], vals...)
}

// Minf returns the minimum of the given values with a given "less" comparator,
// or if no values are given, the zero value of the type.
func Minf[T any](less func(T, T) bool, vals ...T) T {
	if len(vals) == 0 {
		return ZeroValue[T]()
	}

	min := vals[0]
	for _, v := range vals {
		if less(v, min) {
			min = v
		}
	}
	return min
}

// Maxv returns the maximum of the given values, or if no values are given, the zero value of the type.
//
//go:inline
func Maxv[T constraints.Ordered](vals ...T) T {
	return Maxf(Less[T], vals...)
}

// Maxf returns the maximum of the given values with a given "less" comparator,
// or if no values are given, the zero value of the type.
func Maxf[T any](less func(T, T) bool, vals ...T) T {
	if len(vals) == 0 {
		return ZeroValue[T]()
	}

	max := vals[0]
	for _, v := range vals {
		if !less(v, max) {
			max = v
		}
	}
	return max
}

// MinMaxv returns the minimum and maximum of the given values, or if no values are given, two zero value of the type.
//
//go:inline
func MinMaxv[T constraints.Ordered](vals ...T) (T, T) {
	return MinMaxf(Less[T], vals...)
}

// MinMaxv returns the minimum and maximum of the given values with a given "less" comparator,
// or if no values are given, two zero value of the type.
func MinMaxf[T any](less func(T, T) bool, vals ...T) (T, T) {
	if len(vals) == 0 {
		return ZeroValue[T](), ZeroValue[T]()
	}

	min := vals[0]
	max := vals[0]
	for _, v := range vals {
		if less(v, min) {
			min = v
		}
		if !less(v, max) {
			max = v
		}
	}
	return min, max
}

// First returns the first element of the given array, zero value otherwise.
func First[T any](x []T) T {
	if len(x) == 0 {
		return ZeroValue[T]()
	}
	return x[0]
}

// Last returns the last element of the given array, zero value otherwise.
func Last[T any](x []T) T {
	l := len(x)
	if l == 0 {
		return ZeroValue[T]()
	}
	return x[l-1]
}

// ZeroValue returns the zero value of any type.
//
//go:inline
func ZeroValue[T any]() T {
	var t T
	return t
}

// Map calls the function on every array element and returns results in list.
func Map[T, S any](f func(T) S, arr []T) []S {
	what := make([]S, len(arr))
	for i, v := range arr {
		what[i] = f(v)
	}
	return what
}

// Filter calls the function and returns list of values that returned true.
func Filter[T any](f func(T) bool, arr []T) []T {
	what := make([]T, 0, len(arr))
	for _, v := range arr {
		if !f(v) {
			continue
		}
		what = append(what, v)
	}
	return what
}

// RemoveIf calls the function and returns list of values that returned false.
func RemoveIf[T any](f func(T) bool, arr []T) []T {
	what := make([]T, 0, len(arr))
	for _, v := range arr {
		if f(v) {
			continue
		}
		what = append(what, v)
	}
	return what
}

// Take returns up to the first `num` elements.
func Take[T any](num int, arr []T) []T {
	if len(arr) < num {
		return arr
	}
	return arr[:num]
}

// Tail returns up to the last `num` elements.
func Tail[T any](num int, arr []T) []T {
	if len(arr) < num {
		return arr
	}
	return arr[len(arr)-num:]
}

// Drop allocates a new slice, with the first `num` elements dropped.
func Drop[T any, U constraints.Unsigned](num U, arr []T) []T {
	l := U(len(arr))
	if l < num {
		return []T{}
	}

	slice := make([]T, l-num)
	copy(slice, arr[num:])

	return slice
}

// Skips skips the first `num` elements by slicing (underlying array unaffected).
func Skip[T any, U constraints.Unsigned](num U, arr []T) []T {
	if U(len(arr)) < num {
		return arr[:0]
	}
	return arr[num:]
}

// Zip returns a list of tuples, where the i-th tuple contains the i-th element from each of the argument lists.
func Zip[T, U any](a []T, b []U) []Tuple[T, U] {
	l := Min(len(a), len(b))
	what := make([]Tuple[T, U], l)
	for i := 0; i < l; i++ {
		what[i] = NewTuple(a[i], b[i])
	}
	return what
}

// Any returns true if any element in the list matches the given value.
//
//go:inline
func Any[T comparable](val T, arr []T) bool {
	return Anyf(EqualsClosure(val), arr)
}

// Anyf returns true if any elemens in the list returns true when passed to foo.
func Anyf[T any](foo func(v T) bool, arr []T) bool {
	for _, v := range arr {
		if foo(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the list match the given value.
//
//go:inline
func All[T comparable](val T, arr []T) bool {
	return Allf(EqualsClosure(val), arr)
}

// Allf returns true if all elements in the array return true when passed to foo.
func Allf[T any](foo func(v T) bool, arr []T) bool {
	for _, v := range arr {
		if !foo(v) {
			return false
		}
	}
	return true
}

// Repeat creates a list of given size consisting of the same given value.
func Repeat[T any](val T, size int) []T {
	arr := make([]T, size)
	for i := range arr {
		arr[i] = val
	}
	return arr
}

// GetPointer returns the pointer of a given lhs.
func GetPointer[T any](val T) *T {
	return &val
}

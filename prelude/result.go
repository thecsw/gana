package prelude

import "github.com/thecsw/gana"

// Result is a generic result type.
type Result[T, E any] interface {
	// IsOk returns true if the result is an Ok.
	IsOk() bool
	// IsErr returns true if the result is an Err.
	IsErr() bool
	// Unpack returns (value, err) where err is nil if the result is an Ok.
	Unpack() (T, E)
	// UnpackRef returns (&value, &err) where err is nil if the result is an Ok.
	UnpackRef() (*T, *E)
	// Unwrap returns the result's value. Panics if the result is an Err.
	Unwrap() T
	// UnwrapRef returns a reference to the result's value. Panics if the result is an Err.
	UnwrapRef() *T
	// UnwrapErr returns the result's error. Panics if the result is an Ok.
	UnwrapErr() E
	// UnwrapErrRef returns a reference to the result's error. Panics if the result is an Ok.
	UnwrapErrRef() *E
	// UnwrapOr returns the result's value if it is an Ok, otherwise returns the provided default value.
	UnwrapOr(defaultValue T) T
	// UnwrapOrElse returns the result's value if it is an Ok, otherwise returns the result of the provided function.
	UnwrapOrElse(f func() T) T
}

type OkT[T, E any] struct{ value T }
type ErrT[T, E any] struct{ err E }

// Make sure OkT and ErrT implement Result.
var _ Result[int, string] = (*OkT[int, string])(nil)
var _ Result[int, string] = (*ErrT[int, string])(nil)

// Ok returns a new Ok Result with the provided value.
func Ok[T, E any](value T) Result[T, E] { return &OkT[T, E]{value} }

// Err returns a new Err Result with the provided error.
func Err[T, E any](err E) Result[T, E] { return &ErrT[T, E]{err} }

// Try converts (t, err) to a Result.
func Try[T any](t T, err error) Result[T, error] {
	if err == nil {
		return Ok[T, error](t)
	}
	return Err[T](err)
}

func (r *OkT[T, E]) IsOk() bool  { return true }
func (r *OkT[T, E]) IsErr() bool { return false }

func (r *OkT[T, E]) Unpack() (T, E)      { return r.value, gana.ZeroValue[E]() }
func (r *OkT[T, E]) UnpackRef() (*T, *E) { return &r.value, nil }

func (r *OkT[T, E]) Unwrap() T     { return r.value }
func (r *OkT[T, E]) UnwrapRef() *T { return &r.value }

func (r *OkT[T, E]) UnwrapErr() E     { panic("attempted to UnwrapErr an Ok") }
func (r *OkT[T, E]) UnwrapErrRef() *E { panic("attempted to UnwrapErrRef an Ok") }

func (r *OkT[T, E]) UnwrapOr(defaultValue T) T { return r.value }
func (r *OkT[T, E]) UnwrapOrElse(f func() T) T { return r.value }

func (r *ErrT[T, E]) IsOk() bool  { return false }
func (r *ErrT[T, E]) IsErr() bool { return true }

func (r *ErrT[T, E]) Unpack() (T, E)      { return gana.ZeroValue[T](), r.err }
func (r *ErrT[T, E]) UnpackRef() (*T, *E) { return nil, &r.err }

func (r *ErrT[T, E]) Unwrap() T     { panic("attempted to Unwrap an Err") }
func (r *ErrT[T, E]) UnwrapRef() *T { panic("attempted to UnwrapRef an Err") }

func (r *ErrT[T, E]) UnwrapErr() E     { return r.err }
func (r *ErrT[T, E]) UnwrapErrRef() *E { return &r.err }

func (r *ErrT[T, E]) UnwrapOr(defaultValue T) T { return defaultValue }
func (r *ErrT[T, E]) UnwrapOrElse(f func() T) T { return f() }

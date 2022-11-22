package option

// Option is a generic option type.
type Option[T any] interface {
	// IsSome returns true if the option is a Some.
	IsSome() bool
	// IsNone returns true if the option is a None.
	IsNone() bool
	// Unwrap returns the option's value. Panics if the option is a None.
	Unwrap() T
	// UnwrapRef returns a reference to the option's value. Panics if the option is a None.
	UnwrapRef() *T
	// UnwrapNone asserts that the option is a None. Panics if the option is a Some.
	UnwrapNone()
	// UnwrapOr returns the option's value if it is a Some, otherwise returns the provided default value.
	UnwrapOr(defaultValue T) T
	// UnwrapOrElse returns the option's value if it is a Some, otherwise returns the result of the provided function.
	UnwrapOrElse(f func() T) T
}

type SomeT[T any] struct{ value T }
type NoneT[T any] struct{}

// Make sure SomeT and NoneT implement Option.
var _ Option[int] = (*SomeT[int])(nil)
var _ Option[int] = (*NoneT[int])(nil)

// Some returns a new Some.
func Some[T any](value T) Option[T] { return &SomeT[T]{value} }

// None returns a new None.
func None[T any]() Option[T] { return &NoneT[T]{} }

// Maybe converts (t, ok) to an Option.
func Maybe[T any](t T, ok bool) Option[T] {
	if ok {
		return Some(t)
	}
	return None[T]()
}

func (o *SomeT[T]) IsSome() bool { return true }
func (o *SomeT[T]) IsNone() bool { return false }

func (o *SomeT[T]) Unwrap() T     { return o.value }
func (o *SomeT[T]) UnwrapRef() *T { return &o.value }
func (o *SomeT[T]) UnwrapNone()   { panic("attempted to UnwrapNone a Some") }

func (o *SomeT[T]) UnwrapOr(defaultValue T) T { return o.value }
func (o *SomeT[T]) UnwrapOrElse(f func() T) T { return o.value }

func (o *NoneT[T]) IsSome() bool { return false }
func (o *NoneT[T]) IsNone() bool { return true }

func (o *NoneT[T]) Unwrap() T     { panic("attempted to Unwrap a None") }
func (o *NoneT[T]) UnwrapRef() *T { panic("attempted to UnwrapRef a None") }
func (o *NoneT[T]) UnwrapNone()   {}

func (o *NoneT[T]) UnwrapOr(defaultValue T) T { return defaultValue }
func (o *NoneT[T]) UnwrapOrElse(f func() T) T { return f() }

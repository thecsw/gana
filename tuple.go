package gana

// Tuple is a generic tuple type.
type Tuple[A, B any] struct {
	First  A
	Second B
}

// NewTuple returns a new tuple.
func NewTuple[A, B any](first A, second B) Tuple[A, B] {
	return Tuple[A, B]{First: first, Second: second}
}

// Unpack returns the tuple's elements.
func (t Tuple[A, B]) Unpack() (A, B) {
	return t.First, t.Second
}

// UnpackRef returns references to the tuple's elements.
func (t Tuple[A, B]) UnpackRef() (*A, *B) {
	return &t.First, &t.Second
}

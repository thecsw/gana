package gana

// Tuple is a generic tuple type.
type Tuple[A, B any] struct {
	First  A
	Second B
}

// Unpack returns the tuple's elements.
func (t *Tuple[A, B]) Unpack() (A, B) {
	return t.First, t.Second
}

// UnpackRef returns references to the tuple's elements.
func (t *Tuple[A, B]) UnpackRef() (*A, *B) {
	return &t.First, &t.Second
}

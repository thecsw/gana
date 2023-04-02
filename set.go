package gana

type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet returns a new set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

// Add adds a value to the set.
func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

// Remove removes a value from the set.
func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

// Contains returns true if the set contains the value.
func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

// Len returns the number of elements in the set.
func (s *Set[_]) Len() int {
	return len(s.m)
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

// Values returns all values in the set.
func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.m))
	for v := range s.m {
		values = append(values, v)
	}
	return values
}

// Union returns a new set with elements from both s and other.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := NewSet[T]()
	for v := range s.m {
		union.Add(v)
	}
	for v := range other.m {
		union.Add(v)
	}
	return union
}

// Intersection returns a new set with elements that are both in s and other.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersection := NewSet[T]()
	for v := range s.m {
		if other.Contains(v) {
			intersection.Add(v)
		}
	}
	return intersection
}

// Difference returns a new set with elements that are in s but not in other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	difference := NewSet[T]()
	for v := range s.m {
		if !other.Contains(v) {
			difference.Add(v)
		}
	}
	return difference
}

// SymmetricDifference returns a new set with elements that are in either s or other but not both.
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	symmetricDifference := NewSet[T]()
	for v := range s.m {
		if !other.Contains(v) {
			symmetricDifference.Add(v)
		}
	}
	for v := range other.m {
		if !s.Contains(v) {
			symmetricDifference.Add(v)
		}
	}
	return symmetricDifference
}

// IsSubset returns true if s is a subset of other.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for v := range s.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if s is a superset of other.
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

// IsEqual returns true if s and other contain the same elements.
func (s *Set[T]) IsEqual(other *Set[T]) bool {
	return s.IsSubset(other) && other.IsSubset(s)
}

// IsDisjoint returns true if s and other have no elements in common.
func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	for v := range s.m {
		if other.Contains(v) {
			return false
		}
	}
	return true
}

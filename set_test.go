package gana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSet tests the set data structure.
func TestSet(t *testing.T) {
	s := NewSet[int]()
	assert.Equal(t, 0, s.Len())
	s.Add(1)
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Contains(1))
	s.Remove(1)
	assert.Equal(t, 0, s.Len())
	assert.False(t, s.Contains(1))
	s.Add(1)
	s.Add(2)
	s.Add(3)
	assert.Equal(t, 3, s.Len())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
	assert.ElementsMatch(t, []int{1, 2, 3}, s.Values())
	s.Clear()
	assert.Equal(t, 0, s.Len())
	assert.False(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains(3))
	s.Add(1)
	s.Add(2)
	s.Add(3)

	t1 := NewSet[int]()
	t1.Add(1)
	t1.Add(2)
	t2 := NewSet[int]()
	t2.Add(2)
	t2.Add(3)
	assert.ElementsMatch(t, []int{1, 2, 3}, t1.Union(t2).Values())
	assert.ElementsMatch(t, []int{2}, t1.Intersection(t2).Values())
	assert.ElementsMatch(t, []int{1}, t1.Difference(t2).Values())
	assert.ElementsMatch(t, []int{1, 3}, t1.SymmetricDifference(t2).Values())
	assert.True(t, t1.IsSubset(s))
	assert.False(t, t1.IsSuperset(s))
	t1.Add(3)
	assert.True(t, s.IsEqual(t1))
	assert.False(t, s.IsSubset(t2))
	assert.True(t, s.IsSuperset(t2))
	assert.False(t, s.IsEqual(t2))

	t1.Remove(3)
	t2.Remove(2)
	assert.True(t, t1.IsDisjoint(t2))
}

// BenchmarkSet benchmarks the set data structure.
func BenchmarkSet(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

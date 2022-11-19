package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	iter := NewIter([][]int{{1, 2}, {3, 4}})
	result := Collect(Flatten(iter))
	assert.Equal(t, []int{1, 2, 3, 4}, result)

	iter = NewIter([][]int{{1, 2}, {3, 4}})
	s, ok := iter.Next()
	assert.Equal(t, []int{1, 2}, s)
	assert.True(t, ok)
	result = Collect(Flatten(iter))
	assert.Equal(t, []int{3, 4}, result)
}

func TestFlattenIters(t *testing.T) {
	iter := NewIter([]Iter[int]{NewIter([]int{1, 2}), NewIter([]int{3, 4})})
	result := Collect(FlattenIters(iter))
	assert.Equal(t, []int{1, 2, 3, 4}, result)

	iter = NewIter([]Iter[int]{NewIter([]int{1, 2}), NewIter([]int{3, 4})})
	s, ok := iter.Next()
	assert.Equal(t, NewIter([]int{1, 2}), s)
	assert.True(t, ok)
	result = Collect(FlattenIters(iter))
	assert.Equal(t, []int{3, 4}, result)
}

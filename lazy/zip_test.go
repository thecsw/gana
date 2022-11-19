package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	iter := Zip(NewIter([]int{1, 2, 3}), NewIter([]string{"a", "b", "c"}))
	tuple, ok := iter.Next()
	assert.True(t, ok)
	a, b := tuple.Unpack()
	assert.Equal(t, 1, a)
	assert.Equal(t, "a", b)

	tuple, ok = iter.Next()
	assert.True(t, ok)
	a, b = tuple.Unpack()
	assert.Equal(t, 2, a)
	assert.Equal(t, "b", b)

	tuple, ok = iter.Next()
	assert.True(t, ok)
	a, b = tuple.Unpack()
	assert.Equal(t, 3, a)
	assert.Equal(t, "c", b)

	_, ok = iter.Next()
	assert.False(t, ok)

	iter = Zip(NewIter([]int{1, 2, 3}), NewIter([]string{"a", "b"}))
	tuple, ok = iter.Next()
	assert.True(t, ok)
	a, b = tuple.Unpack()
	assert.Equal(t, 1, a)
	assert.Equal(t, "a", b)

	tuple, ok = iter.Next()
	assert.True(t, ok)
	a, b = tuple.Unpack()
	assert.Equal(t, 2, a)
	assert.Equal(t, "b", b)

	_, ok = iter.Next()
	assert.False(t, ok)
}

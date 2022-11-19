package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	result := Collect(Take(iter, 2))
	assert.Equal(t, []string{"a", "b"}, result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = Collect(Take(iter, 2))
	assert.Equal(t, []string{"b", "c"}, result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok = iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	s, ok = iter.Next()
	assert.Equal(t, "b", s)
	assert.True(t, ok)
	result = Collect(Take(iter, 2))
	assert.Equal(t, []string{"c"}, result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok = iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	s, ok = iter.Next()
	assert.Equal(t, "b", s)
	assert.True(t, ok)
	s, ok = iter.Next()
	assert.Equal(t, "c", s)
	assert.True(t, ok)
	result = Collect(Take(iter, 2))
	assert.Equal(t, []string{}, result)

	iter = NewIter([]string{"a", "b", "c"})
	result = Collect(Take(iter, 0))
	assert.Equal(t, []string{}, result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok = iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = Collect(Take(iter, 0))
	assert.Equal(t, []string{}, result)
}

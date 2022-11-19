package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkip(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	result := Collect(Skip(iter, 2))
	assert.Equal(t, []string{"c"}, result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = Collect(Skip(iter, 2))
	assert.Equal(t, []string{}, result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok = iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	s, ok = iter.Next()
	assert.Equal(t, "b", s)
	assert.True(t, ok)
	result = Collect(Skip(iter, 2))
	assert.Equal(t, []string{}, result)
}

func TestSkipWhile(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	result := Collect(SkipWhile(iter, func(s string) bool { return s != "b" }))
	assert.Equal(t, []string{"b", "c"}, result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = Collect(SkipWhile(iter, func(s string) bool { return s == "b" }))
	assert.Equal(t, []string{"c"}, result)
}

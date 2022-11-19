package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	result := Collect(Map(iter, func(s string) string { return s + s }))
	assert.Equal(t, []string{"aa", "bb", "cc"}, result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = Collect(Map(iter, func(s string) string { return s + s }))
	assert.Equal(t, []string{"bb", "cc"}, result)
}

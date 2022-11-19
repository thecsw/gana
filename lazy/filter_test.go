package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	result := Collect(Filter(iter, func(s string) bool { return s != "b" }))
	assert.Equal(t, []string{"a", "c"}, result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = Collect(Filter(iter, func(s string) bool { return s != "b" }))
	assert.Equal(t, []string{"c"}, result)
}

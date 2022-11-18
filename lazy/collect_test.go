package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollect(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	result := Collect(iter)
	assert.Equal(t, []string{"a", "b", "c"}, result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = Collect(iter)
	assert.Equal(t, []string{"b", "c"}, result)

	iter = NewIter([]string{})
	result = Collect(iter)
	assert.Equal(t, []string{}, result)
}

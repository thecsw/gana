package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thecsw/gana"
)

func TestIter(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	s, ok = iter.Next()
	assert.Equal(t, "b", s)
	assert.True(t, ok)
	s, ok = iter.Next()
	assert.Equal(t, "c", s)
	assert.True(t, ok)
	s, ok = iter.Next()
	assert.Equal(t, gana.ZeroValue[string](), s)
	assert.False(t, ok)
}

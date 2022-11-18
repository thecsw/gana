package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thecsw/gana"
)

func TestFold(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	result := Fold(iter, "", func(a, b string) string { return a + b })
	assert.Equal(t, "abc", result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = Fold(iter, "", func(a, b string) string { return a + b })
	assert.Equal(t, "bc", result)
}

func TestFold1(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	result := Fold1(iter, func(a, b string) string { return a + b })
	assert.Equal(t, "abc", result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = Fold1(iter, func(a, b string) string { return a + b })
	assert.Equal(t, "bc", result)
}

func TestFold1Empty(t *testing.T) {
	iter := NewIter([]string{})
	assert.Equal(t, gana.ZeroValue[string](), Fold1(iter, func(a, b string) string { return a + b }))
}

func TestFoldRight(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	result := FoldRight(iter, "", func(b, a string) string { return a + b })
	assert.Equal(t, "cba", result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = FoldRight(iter, "", func(b, a string) string { return a + b })
	assert.Equal(t, "cb", result)
}

func TestFoldRight1(t *testing.T) {
	iter := NewIter([]string{"a", "b", "c"})
	result := FoldRight1(iter, func(b, a string) string { return a + b })
	assert.Equal(t, "cba", result)

	iter = NewIter([]string{"a", "b", "c"})
	s, ok := iter.Next()
	assert.Equal(t, "a", s)
	assert.True(t, ok)
	result = FoldRight1(iter, func(b, a string) string { return a + b })
	assert.Equal(t, "cb", result)
}

func TestFoldRight1Empty(t *testing.T) {
	iter := NewIter([]string{})
	assert.Equal(t, gana.ZeroValue[string](), FoldRight1(iter, func(b, a string) string { return a + b }))
}

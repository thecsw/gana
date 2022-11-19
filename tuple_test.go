package gana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTuple(t *testing.T) {
	tuple := NewTuple(1, 2)
	a, b := tuple.Unpack()
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, b)
}

func TestUnpack(t *testing.T) {
	a, b := Tuple[int, int]{1, 2}.Unpack()
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, b)
}

func TestUnpackRef(t *testing.T) {
	a, b := Tuple[int, int]{1, 2}.UnpackRef()
	assert.Equal(t, 1, *a)
	assert.Equal(t, 2, *b)
}

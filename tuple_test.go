package gana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	a, b := Tuple[int, int]{1, 2}.Unpack()
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, b)
}

func TestUnpackRef(t *testing.T) {
	a, b := (&Tuple[int, int]{1, 2}).UnpackRef()
	assert.Equal(t, 1, *a)
	assert.Equal(t, 2, *b)
}

package prelude

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSome(t *testing.T) {
	assert.True(t, Some(1).IsSome())
	assert.False(t, None[int]().IsSome())
}

func TestIsNone(t *testing.T) {
	assert.False(t, Some(1).IsNone())
	assert.True(t, None[int]().IsNone())
}

func TestUnpack(t *testing.T) {
	v, ok := Some(1).Unpack()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = None[int]().Unpack()
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func TestUnpackRef(t *testing.T) {
	v, ok := Some(1).UnpackRef()
	assert.True(t, ok)
	assert.Equal(t, 1, *v)

	v, ok = None[int]().UnpackRef()
	assert.False(t, ok)
	assert.Nil(t, v)
}

func TestUnwrap(t *testing.T) {
	assert.Equal(t, 1, Some(1).Unwrap())
	assert.Panics(t, func() { None[int]().Unwrap() })
}

func TestUnwrapRef(t *testing.T) {
	assert.Equal(t, 1, *Some(1).UnwrapRef())
	assert.Panics(t, func() { None[int]().UnwrapRef() })
}

func TestUnwrapNone(t *testing.T) {
	assert.Panics(t, func() { Some(1).UnwrapNone() })
	assert.NotPanics(t, func() { None[int]().UnwrapNone() })
}

func TestUnwrapOr(t *testing.T) {
	assert.Equal(t, 1, Some(1).UnwrapOr(2))
	assert.Equal(t, 2, None[int]().UnwrapOr(2))
}

func TestUnwrapOrElse(t *testing.T) {
	assert.Equal(t, 1, Some(1).UnwrapOrElse(func() int { return 2 }))
	assert.Equal(t, 2, None[int]().UnwrapOrElse(func() int { return 2 }))
}

func TestMaybe(t *testing.T) {
	o := Maybe(1, true)
	assert.True(t, o.IsSome())
	assert.Equal(t, 1, o.Unwrap())

	o = Maybe(1, false)
	assert.True(t, o.IsNone())
}

func TestOr(t *testing.T) {
	o := Some(1).Or(Some(2))
	assert.True(t, o.IsSome())
	assert.Equal(t, 1, o.Unwrap())

	o = None[int]().Or(Some(2))
	assert.True(t, o.IsSome())
	assert.Equal(t, 2, o.Unwrap())
}

func TestXor(t *testing.T) {
	o := Some(1).Xor(Some(2))
	assert.True(t, o.IsNone())

	o = Some(1).Xor(None[int]())
	assert.True(t, o.IsSome())
	assert.Equal(t, 1, o.Unwrap())

	o = None[int]().Xor(Some(2))
	assert.True(t, o.IsSome())
	assert.Equal(t, 2, o.Unwrap())

	o = None[int]().Xor(None[int]())
	assert.True(t, o.IsNone())
}

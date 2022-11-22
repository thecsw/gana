package prelude

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thecsw/gana"
)

func TestIsOk(t *testing.T) {
	assert.True(t, Ok[int, int](1).IsOk())
	assert.False(t, Err[int](1).IsOk())
}

func TestIsErr(t *testing.T) {
	assert.False(t, Ok[int, int](1).IsErr())
	assert.True(t, Err[int](1).IsErr())
}

func TestUnpackResult(t *testing.T) {
	v, err := Ok[int, error](1).Unpack()
	assert.Nil(t, err)
	assert.Equal(t, 1, v)

	v, err = Err[int](errors.New("test error")).Unpack()
	assert.NotNil(t, err)
	assert.Equal(t, gana.ZeroValue[int](), v)
}

func TestUnpackRefResult(t *testing.T) {
	v, err := Ok[int, error](1).UnpackRef()
	assert.Nil(t, err)
	assert.Equal(t, 1, *v)

	v, err = Err[int](errors.New("test error")).UnpackRef()
	assert.NotNil(t, err)
	assert.Nil(t, v)
}

func TestUnwrapResult(t *testing.T) {
	assert.Equal(t, 1, Ok[int, int](1).Unwrap())
	assert.Panics(t, func() { Err[int](1).Unwrap() })
}

func TestUnwrapRefResult(t *testing.T) {
	assert.Equal(t, 1, *Ok[int, int](1).UnwrapRef())
	assert.Panics(t, func() { Err[int](1).UnwrapRef() })
}

func TestUnwrapErrResult(t *testing.T) {
	assert.Panics(t, func() { Ok[int, int](1).UnwrapErr() })

	e := Err[int](1)
	assert.Equal(t, 1, e.UnwrapErr())
	assert.NotPanics(t, func() { e.UnwrapErr() })
}

func TestUnwrapErrRefResult(t *testing.T) {
	assert.Panics(t, func() { Ok[int, int](1).UnwrapErrRef() })

	e := Err[int](1)
	assert.Equal(t, 1, *e.UnwrapErrRef())
	assert.NotPanics(t, func() { e.UnwrapErrRef() })
}

func TestUnwrapOrResult(t *testing.T) {
	assert.Equal(t, 1, Ok[int, int](1).UnwrapOr(2))
	assert.Equal(t, 2, Err[int](1).UnwrapOr(2))
}

func TestUnwrapOrElseResult(t *testing.T) {
	assert.Equal(t, 1, Ok[int, int](1).UnwrapOrElse(func() int { return 2 }))
	assert.Equal(t, 2, Err[int](1).UnwrapOrElse(func() int { return 2 }))
}

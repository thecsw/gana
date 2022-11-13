package gana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipRunes(t *testing.T) {
	assert.Equal(t, SkipRunes(1, "abc"), "bc", `"abc"[1:]`)
	assert.Equal(t, SkipRunes(2, "abc"), "c", `"abc"[2:]`)
	assert.Equal(t, SkipRunes(3, "abc"), "", `"abc"[3:]`)
	assert.Equal(t, SkipRunes(4, "abc"), "", `"abc"[4:]`)
	assert.Equal(t, SkipRunes(0, "abc"), "abc", `"abc"[0:]`)
	assert.Equal(t, SkipRunes(1, ""), "", `""[1:]`)
	// Emojis
	assert.Equal(t, SkipRunes(1, "a😀"), "😀", `"a😀"[1:]`)
	assert.Equal(t, SkipRunes(2, "a😀"), "", `"a😀"[2:]`)
	assert.Equal(t, SkipRunes(3, "a😀"), "", `"a😀"[3:]`)
	assert.Equal(t, SkipRunes(4, "a😀"), "", `"a😀"[4:]`)
	assert.Equal(t, SkipRunes(0, "a😀"), "a😀", `"a😀"[0:]`)
	// Complex emoji
	assert.Equal(t, SkipRunes(1, "a😀😀"), "😀😀", `"a😀😀"[1:]`)
	assert.Equal(t, SkipRunes(2, "a😀😀"), "😀", `"a😀😀"[2:]`)
	assert.Equal(t, SkipRunes(3, "a😀😀"), "", `"a😀😀"[3:]`)
	assert.Equal(t, SkipRunes(4, "a😀😀"), "", `"a😀😀"[4:]`)
	assert.Equal(t, SkipRunes(0, "a😀😀"), "a😀😀", `"a😀😀"[0:]`)
}

func TestDropRunes(t *testing.T) {
	assert.Equal(t, DropRunes(1, "abc"), "bc", `"abc"[1:]`)
	assert.Equal(t, DropRunes(2, "abc"), "c", `"abc"[2:]`)
	assert.Equal(t, DropRunes(3, "abc"), "", `"abc"[3:]`)
	assert.Equal(t, DropRunes(4, "abc"), "", `"abc"[4:]`)
	assert.Equal(t, DropRunes(0, "abc"), "abc", `"abc"[0:]`)
	assert.Equal(t, DropRunes(1, ""), "", `""[1:]`)
	// Emojis
	assert.Equal(t, DropRunes(1, "a😀"), "😀", `"a😀"[1:]`)
	assert.Equal(t, DropRunes(2, "a😀"), "", `"a😀"[2:]`)
	assert.Equal(t, DropRunes(3, "a😀"), "", `"a😀"[3:]`)
	assert.Equal(t, DropRunes(4, "a😀"), "", `"a😀"[4:]`)
	assert.Equal(t, DropRunes(0, "a😀"), "a😀", `"a😀"[0:]`)
	// Complex emoji
	assert.Equal(t, DropRunes(1, "a😀😀"), "😀😀", `"a😀😀"[1:]`)
	assert.Equal(t, DropRunes(2, "a😀😀"), "😀", `"a😀😀"[2:]`)
	assert.Equal(t, DropRunes(3, "a😀😀"), "", `"a😀😀"[3:]`)
	assert.Equal(t, DropRunes(4, "a😀😀"), "", `"a😀😀"[4:]`)
	assert.Equal(t, DropRunes(0, "a😀😀"), "a😀😀", `"a😀😀"[0:]`)
}

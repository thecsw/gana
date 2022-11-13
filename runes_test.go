package gana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipRunes(t *testing.T) {
	assert.Equal(t, SkipRunes[uint](1, "abc"), "bc", `"abc"[1:]`)
	assert.Equal(t, SkipRunes[uint](2, "abc"), "c", `"abc"[2:]`)
	assert.Equal(t, SkipRunes[uint](3, "abc"), "", `"abc"[3:]`)
	assert.Equal(t, SkipRunes[uint](4, "abc"), "", `"abc"[4:]`)
	assert.Equal(t, SkipRunes[uint](0, "abc"), "abc", `"abc"[0:]`)
	assert.Equal(t, SkipRunes[uint](1, ""), "", `""[1:]`)
	// Emojis
	assert.Equal(t, SkipRunes[uint](1, "a😀"), "😀", `"a😀"[1:]`)
	assert.Equal(t, SkipRunes[uint](2, "a😀"), "", `"a😀"[2:]`)
	assert.Equal(t, SkipRunes[uint](3, "a😀"), "", `"a😀"[3:]`)
	assert.Equal(t, SkipRunes[uint](4, "a😀"), "", `"a😀"[4:]`)
	assert.Equal(t, SkipRunes[uint](0, "a😀"), "a😀", `"a😀"[0:]`)
	// Complex emoji
	assert.Equal(t, SkipRunes[uint](1, "a😀😀"), "😀😀", `"a😀😀"[1:]`)
	assert.Equal(t, SkipRunes[uint](2, "a😀😀"), "😀", `"a😀😀"[2:]`)
	assert.Equal(t, SkipRunes[uint](3, "a😀😀"), "", `"a😀😀"[3:]`)
	assert.Equal(t, SkipRunes[uint](4, "a😀😀"), "", `"a😀😀"[4:]`)
	assert.Equal(t, SkipRunes[uint](0, "a😀😀"), "a😀😀", `"a😀😀"[0:]`)
}

func TestDropRunes(t *testing.T) {
	assert.Equal(t, DropRunes[uint](1, "abc"), "bc", `"abc"[1:]`)
	assert.Equal(t, DropRunes[uint](2, "abc"), "c", `"abc"[2:]`)
	assert.Equal(t, DropRunes[uint](3, "abc"), "", `"abc"[3:]`)
	assert.Equal(t, DropRunes[uint](4, "abc"), "", `"abc"[4:]`)
	assert.Equal(t, DropRunes[uint](0, "abc"), "abc", `"abc"[0:]`)
	assert.Equal(t, DropRunes[uint](1, ""), "", `""[1:]`)
	// Emojis
	assert.Equal(t, DropRunes[uint](1, "a😀"), "😀", `"a😀"[1:]`)
	assert.Equal(t, DropRunes[uint](2, "a😀"), "", `"a😀"[2:]`)
	assert.Equal(t, DropRunes[uint](3, "a😀"), "", `"a😀"[3:]`)
	assert.Equal(t, DropRunes[uint](4, "a😀"), "", `"a😀"[4:]`)
	assert.Equal(t, DropRunes[uint](0, "a😀"), "a😀", `"a😀"[0:]`)
	// Complex emoji
	assert.Equal(t, DropRunes[uint](1, "a😀😀"), "😀😀", `"a😀😀"[1:]`)
	assert.Equal(t, DropRunes[uint](2, "a😀😀"), "😀", `"a😀😀"[2:]`)
	assert.Equal(t, DropRunes[uint](3, "a😀😀"), "", `"a😀😀"[3:]`)
	assert.Equal(t, DropRunes[uint](4, "a😀😀"), "", `"a😀😀"[4:]`)
	assert.Equal(t, DropRunes[uint](0, "a😀😀"), "a😀😀", `"a😀😀"[0:]`)
}

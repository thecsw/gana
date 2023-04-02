package gana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipRunes(t *testing.T) {
	assert.Equal(t, "bc", SkipRunes[uint](1, "abc"), `"abc"[1:]`)
	assert.Equal(t, "c", SkipRunes[uint](2, "abc"), `"abc"[2:]`)
	assert.Equal(t, "", SkipRunes[uint](3, "abc"), `"abc"[3:]`)
	assert.Equal(t, "", SkipRunes[uint](4, "abc"), `"abc"[4:]`)
	assert.Equal(t, "abc", SkipRunes[uint](0, "abc"), `"abc"[0:]`)
	assert.Equal(t, "", SkipRunes[uint](1, ""), `""[1:]`)
	// Emojis
	assert.Equal(t, "😀", SkipRunes[uint](1, "a😀"), `"a😀"[1:]`)
	assert.Equal(t, "", SkipRunes[uint](2, "a😀"), `"a😀"[2:]`)
	assert.Equal(t, "", SkipRunes[uint](3, "a😀"), `"a😀"[3:]`)
	assert.Equal(t, "", SkipRunes[uint](4, "a😀"), `"a😀"[4:]`)
	assert.Equal(t, "a😀", SkipRunes[uint](0, "a😀"), `"a😀"[0:]`)
	// Complex emoji
	assert.Equal(t, "😀😀", SkipRunes[uint](1, "a😀😀"), `"a😀😀"[1:]`)
	assert.Equal(t, "😀", SkipRunes[uint](2, "a😀😀"), `"a😀😀"[2:]`)
	assert.Equal(t, "", SkipRunes[uint](3, "a😀😀"), `"a😀😀"[3:]`)
	assert.Equal(t, "", SkipRunes[uint](4, "a😀😀"), `"a😀😀"[4:]`)
	assert.Equal(t, "a😀😀", SkipRunes[uint](0, "a😀😀"), `"a😀😀"[0:]`)
}

func TestDropRunes(t *testing.T) {
	assert.Equal(t, "bc", DropRunes[uint](1, "abc"), `"abc"[1:]`)
	assert.Equal(t, "c", DropRunes[uint](2, "abc"), `"abc"[2:]`)
	assert.Equal(t, "", DropRunes[uint](3, "abc"), `"abc"[3:]`)
	assert.Equal(t, "", DropRunes[uint](4, "abc"), `"abc"[4:]`)
	assert.Equal(t, "abc", DropRunes[uint](0, "abc"), `"abc"[0:]`)
	assert.Equal(t, "", DropRunes[uint](1, ""), `""[1:]`)
	// Emojis
	assert.Equal(t, "😀", DropRunes[uint](1, "a😀"), `"a😀"[1:]`)
	assert.Equal(t, "", DropRunes[uint](2, "a😀"), `"a😀"[2:]`)
	assert.Equal(t, "", DropRunes[uint](3, "a😀"), `"a😀"[3:]`)
	assert.Equal(t, "", DropRunes[uint](4, "a😀"), `"a😀"[4:]`)
	assert.Equal(t, "a😀", DropRunes[uint](0, "a😀"), `"a😀"[0:]`)
	// Complex emoji
	assert.Equal(t, "😀😀", DropRunes[uint](1, "a😀😀"), `"a😀😀"[1:]`)
	assert.Equal(t, "😀", DropRunes[uint](2, "a😀😀"), `"a😀😀"[2:]`)
	assert.Equal(t, "", DropRunes[uint](3, "a😀😀"), `"a😀😀"[3:]`)
	assert.Equal(t, "", DropRunes[uint](4, "a😀😀"), `"a😀😀"[4:]`)
	assert.Equal(t, "a😀😀", DropRunes[uint](0, "a😀😀"), `"a😀😀"[0:]`)
}

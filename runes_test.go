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
	assert.Equal(t, SkipRunes[uint](1, "ağ"), "ğ", `"ağ"[1:]`)
	assert.Equal(t, SkipRunes[uint](2, "ağ"), "", `"ağ"[2:]`)
	assert.Equal(t, SkipRunes[uint](3, "ağ"), "", `"ağ"[3:]`)
	assert.Equal(t, SkipRunes[uint](4, "ağ"), "", `"ağ"[4:]`)
	assert.Equal(t, SkipRunes[uint](0, "ağ"), "ağ", `"ağ"[0:]`)
	// Complex emoji
	assert.Equal(t, SkipRunes[uint](1, "ağğ"), "ğğ", `"ağğ"[1:]`)
	assert.Equal(t, SkipRunes[uint](2, "ağğ"), "ğ", `"ağğ"[2:]`)
	assert.Equal(t, SkipRunes[uint](3, "ağğ"), "", `"ağğ"[3:]`)
	assert.Equal(t, SkipRunes[uint](4, "ağğ"), "", `"ağğ"[4:]`)
	assert.Equal(t, SkipRunes[uint](0, "ağğ"), "ağğ", `"ağğ"[0:]`)
}

func TestDropRunes(t *testing.T) {
	assert.Equal(t, DropRunes[uint](1, "abc"), "bc", `"abc"[1:]`)
	assert.Equal(t, DropRunes[uint](2, "abc"), "c", `"abc"[2:]`)
	assert.Equal(t, DropRunes[uint](3, "abc"), "", `"abc"[3:]`)
	assert.Equal(t, DropRunes[uint](4, "abc"), "", `"abc"[4:]`)
	assert.Equal(t, DropRunes[uint](0, "abc"), "abc", `"abc"[0:]`)
	assert.Equal(t, DropRunes[uint](1, ""), "", `""[1:]`)
	// Emojis
	assert.Equal(t, DropRunes[uint](1, "ağ"), "ğ", `"ağ"[1:]`)
	assert.Equal(t, DropRunes[uint](2, "ağ"), "", `"ağ"[2:]`)
	assert.Equal(t, DropRunes[uint](3, "ağ"), "", `"ağ"[3:]`)
	assert.Equal(t, DropRunes[uint](4, "ağ"), "", `"ağ"[4:]`)
	assert.Equal(t, DropRunes[uint](0, "ağ"), "ağ", `"ağ"[0:]`)
	// Complex emoji
	assert.Equal(t, DropRunes[uint](1, "ağğ"), "ğğ", `"ağğ"[1:]`)
	assert.Equal(t, DropRunes[uint](2, "ağğ"), "ğ", `"ağğ"[2:]`)
	assert.Equal(t, DropRunes[uint](3, "ağğ"), "", `"ağğ"[3:]`)
	assert.Equal(t, DropRunes[uint](4, "ağğ"), "", `"ağğ"[4:]`)
	assert.Equal(t, DropRunes[uint](0, "ağğ"), "ağğ", `"ağğ"[0:]`)
}

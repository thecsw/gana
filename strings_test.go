package gana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropString(t *testing.T) {
	assert.Equal(t, DropString[uint](1, "abc"), "bc", `"abc"[1:]`)
	assert.Equal(t, DropString[uint](2, "abc"), "c", `"abc"[2:]`)
	assert.Equal(t, DropString[uint](3, "abc"), "", `"abc"[3:]`)
	assert.Equal(t, DropString[uint](4, "abc"), "", `"abc"[4:]`)
	assert.Equal(t, DropString[uint](0, "abc"), "abc", `"abc"[0:]`)
	assert.Equal(t, DropString[uint](1, ""), "", `""[1:]`)
}

func TestSkipString(t *testing.T) {
	assert.Equal(t, SkipString[uint](1, "abc"), "bc", `"abc"[1:]`)
	assert.Equal(t, SkipString[uint](2, "abc"), "c", `"abc"[2:]`)
	assert.Equal(t, SkipString[uint](3, "abc"), "", `"abc"[3:]`)
	assert.Equal(t, SkipString[uint](4, "abc"), "", `"abc"[4:]`)
	assert.Equal(t, SkipString[uint](0, "abc"), "abc", `"abc"[0:]`)
	assert.Equal(t, SkipString[uint](1, ""), "", `""[1:]`)
}

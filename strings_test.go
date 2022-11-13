package gana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropString(t *testing.T) {
	assert.Equal(t, DropString(1, "abc"), "bc", `"abc"[1:]`)
	assert.Equal(t, DropString(2, "abc"), "c", `"abc"[2:]`)
	assert.Equal(t, DropString(3, "abc"), "", `"abc"[3:]`)
	assert.Equal(t, DropString(4, "abc"), "", `"abc"[4:]`)
	assert.Equal(t, DropString(0, "abc"), "abc", `"abc"[0:]`)
	assert.Equal(t, DropString(1, ""), "", `""[1:]`)
}

func TestSkipString(t *testing.T) {
	assert.Equal(t, SkipString(1, "abc"), "bc", `"abc"[1:]`)
	assert.Equal(t, SkipString(2, "abc"), "c", `"abc"[2:]`)
	assert.Equal(t, SkipString(3, "abc"), "", `"abc"[3:]`)
	assert.Equal(t, SkipString(4, "abc"), "", `"abc"[4:]`)
	assert.Equal(t, SkipString(0, "abc"), "abc", `"abc"[0:]`)
	assert.Equal(t, SkipString(1, ""), "", `""[1:]`)
}

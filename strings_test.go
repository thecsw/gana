package gana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeString(t *testing.T) {
	assert.Equal(t, "a", TakeString[uint](1, "abc"), `"abc"[:1]`)
	assert.Equal(t, "ab", TakeString[uint](2, "abc"), `"abc"[:2]`)
	assert.Equal(t, "abc", TakeString[uint](3, "abc"), `"abc"[:3]`)
	assert.Equal(t, "abc", TakeString[uint](4, "abc"), `"abc"[:4]`)
	assert.Equal(t, "", TakeString[uint](0, "abc"), `"abc"[:0]`)
	assert.Equal(t, "", TakeString[uint](1, ""), `""[:1]`)
}

func TestDropString(t *testing.T) {
	assert.Equal(t, "bc", DropString[uint](1, "abc"), `"abc"[1:]`)
	assert.Equal(t, "c", DropString[uint](2, "abc"), `"abc"[2:]`)
	assert.Equal(t, "", DropString[uint](3, "abc"), `"abc"[3:]`)
	assert.Equal(t, "", DropString[uint](4, "abc"), `"abc"[4:]`)
	assert.Equal(t, "abc", DropString[uint](0, "abc"), `"abc"[0:]`)
	assert.Equal(t, "", DropString[uint](1, ""), `""[1:]`)
}

func TestSkipString(t *testing.T) {
	assert.Equal(t, "bc", SkipString[uint](1, "abc"), `"abc"[1:]`)
	assert.Equal(t, "c", SkipString[uint](2, "abc"), `"abc"[2:]`)
	assert.Equal(t, "", SkipString[uint](3, "abc"), `"abc"[3:]`)
	assert.Equal(t, "", SkipString[uint](4, "abc"), `"abc"[4:]`)
	assert.Equal(t, "abc", SkipString[uint](0, "abc"), `"abc"[0:]`)
	assert.Equal(t, "", SkipString[uint](1, ""), `""[1:]`)
}

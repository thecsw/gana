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

func TestPeekString(t *testing.T) {
	assert.Equal(t, "a", PeekString[uint](1, "abc"), `"abc"[:1]`)
	assert.Equal(t, "ab", PeekString[uint](2, "abc"), `"abc"[:2]`)
	assert.Equal(t, "abc", PeekString[uint](3, "abc"), `"abc"[:3]`)
	assert.Equal(t, "abc", PeekString[uint](4, "abc"), `"abc"[:4]`)
	assert.Equal(t, "", PeekString[uint](0, "abc"), `"abc"[:0]`)
	assert.Equal(t, "", PeekString[uint](1, ""), `""[:1]`)
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

func TestCountRunesLeft(t *testing.T) {
	assert.Equal(t, uint(2), CountRunesLeft[uint]("aabc", 'a'), `CountRunesLeft("aabc", 'a')`)
	assert.Equal(t, uint(1), CountRunesLeft[uint]("bc", 'b'), `CountRunesLeft("bc", 'b')`)
	assert.Equal(t, uint(0), CountRunesLeft[uint]("aabc", 'c'), `CountRunesLeft("aabc", 'c')`)
	assert.Equal(t, uint(0), CountRunesLeft[uint]("aabc", 'd'), `CountRunesLeft("aabc", 'd')`)
	assert.Equal(t, uint(0), CountRunesLeft[uint]("", 'a'), `CountRunesLeft("", 'a')`)
}

func TestCountRunesRight(t *testing.T) {
	assert.Equal(t, uint(2), CountRunesRight[uint]("aabcc", 'c'), `CountRunesRight("aabcc", 'c')`)
	assert.Equal(t, uint(1), CountRunesRight[uint]("aab", 'b'), `CountRunesRight("aab", 'b')`)
	assert.Equal(t, uint(0), CountRunesRight[uint]("aabc", 'a'), `CountRunesRight("aabc", 'a')`)
	assert.Equal(t, uint(0), CountRunesRight[uint]("aabc", 'd'), `CountRunesRight("aabc", 'd')`)
	assert.Equal(t, uint(0), CountRunesRight[uint]("", 'a'), `CountRunesRight("", 'a')`)
}

func TestCountStringsLeft(t *testing.T) {
	assert.Equal(t, uint(1), CountStringsLeft[uint]("aabc", "aa"), `CountStringsLeft("aabc", "aa")`)
	assert.Equal(t, uint(3), CountStringsLeft[uint]("bcbcbca", "bc"), `CountStringsLeft("bcbcbca", "b")`)
	assert.Equal(t, uint(0), CountStringsLeft[uint]("aabc", "c"), `CountStringsLeft("aabc", "c")`)
	assert.Equal(t, uint(0), CountStringsLeft[uint]("aabc", "d"), `CountStringsLeft("aabc", "d")`)
	assert.Equal(t, uint(0), CountStringsLeft[uint]("", "a"), `CountStringsLeft("", "a")`)
}

func TestCountStringsRight(t *testing.T) {
	assert.Equal(t, uint(1), CountStringsRight[uint]("aabc", "bc"), `CountStringsRight("aabc", "bc")`)
	assert.Equal(t, uint(3), CountStringsRight[uint]("abcbcbc", "bc"), `CountStringsRight("abcbcbc", "bc")`)
	assert.Equal(t, uint(0), CountStringsRight[uint]("aabc", "a"), `CountStringsRight("aabc", "a")`)
	assert.Equal(t, uint(0), CountStringsRight[uint]("aabc", "d"), `CountStringsRight("aabc", "d")`)
	assert.Equal(t, uint(0), CountStringsRight[uint]("", "a"), `CountStringsRight("", "a")`)
}

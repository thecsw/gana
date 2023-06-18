package gana

import (
	"strings"

	"golang.org/x/exp/constraints"
)

// TakeString allocates a new string, with the first `num` bytes of a string taken.
func TakeString[T constraints.Unsigned](num T, what string) string {
	builder := strings.Builder{}
	builder.WriteString(what[:Min(num, T(len(what)))])
	return builder.String()
}

// PeekString takes the first `num` bytes of a string by slicing (underlying array unaffected).
func PeekString[T constraints.Unsigned](num T, what string) string {
	return what[:Min(num, T(len(what)))]
}

// DropString allocates a new string, with the first `num` bytes of a string dropped.
func DropString[T constraints.Unsigned](num T, what string) string {
	if T(len(what)) < num {
		return ""
	}
	builder := strings.Builder{}
	builder.WriteString(what[num:])
	return builder.String()
}

// SkipString skips the first `num` bytes of a string by slicing (underlying array unaffected).
func SkipString[T constraints.Unsigned](num T, what string) string {
	if T(len(what)) < num {
		return what[:0]
	}
	return what[num:]
}

// CountRunesLeft counts the number of times a rune appears
// at the beginning of a string.
func CountRunesLeft[T constraints.Unsigned](s string, r rune) T {
	count := T(0)
	for _, rr := range s {
		if rr != r {
			return count
		}
		count++
	}
	return count
}

// CountRunesRight counts the number of times a rune appears
// at the end of a string.
func CountRunesRight[T constraints.Unsigned](s string, r rune) T {
	count := T(0)
	for i := len(s) - 1; i >= 0; i-- {
		if rune(s[i]) != r {
			return count
		}
		count++
	}
	return count
}

// CountStringsLeft counts the number of times a substring appears
// at the beginning of a string.
func CountStringsLeft[T constraints.Unsigned](s, substr string) T {
	count := T(0)
	for strings.HasPrefix(s, substr) {
		count++
		s = strings.TrimPrefix(s, substr)
	}
	return count
}

// CountStringsRight counts the number of times a substring appears
// at the end of a string.
func CountStringsRight[T constraints.Unsigned](s, substr string) T {
	count := T(0)
	for strings.HasSuffix(s, substr) {
		count++
		s = strings.TrimSuffix(s, substr)
	}
	return count
}

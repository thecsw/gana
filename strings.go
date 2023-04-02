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

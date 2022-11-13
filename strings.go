package gana

import "strings"

// DropString allocates a new string, with the first `num` bytes of a string dropped.
func DropString(num uint, what string) string {
	l := uint(len(what))
	if l < num {
		return ""
	}

	builder := strings.Builder{}
	builder.WriteString(what[num:])

	return builder.String()
}

// SkipString skips the first `num` bytes of a string by slicing (underlying array unaffected).
func SkipString(num uint, what string) string {
	if uint(len(what)) < num {
		return what[:0]
	}
	return what[num:]
}

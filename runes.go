package gana

import "golang.org/x/exp/constraints"

// SkipRunes skips the first `num` runes of a string by slicing (underlying array unaffected).
func SkipRunes[T constraints.Unsigned](num T, what string) string {
	if T(len(what)) < num {
		return ""
	}
	for i := range what {
		if num == 0 {
			return SkipString(T(i), what)
		}
		num--
	}
	// Unreachable
	return ""
}

// DropRunes allocates a new string, with the first `num` runes of a string dropped.
func DropRunes[T constraints.Unsigned](num T, what string) string {
	if T(len(what)) < num {
		return ""
	}
	for i := range what {
		if num == 0 {
			return DropString(T(i), what)
		}
		num--
	}
	// Unreachable
	return ""
}

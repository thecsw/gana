package gana

// SkipRunes skips the first `num` runes of a string by slicing (underlying array unaffected).
func SkipRunes(num int, what string) string {
	if len(what) < num {
		return ""
	}
	if num < 0 {
		return what
	}
	for i := range what {
		if num == 0 {
			return SkipString(i, what)
		}
		num--
	}
	// Unreachable
	return ""
}

// DropRunes allocates a new string, with the first `num` runes of a string dropped.
func DropRunes(num int, what string) string {
	if len(what) < num {
		return ""
	}
	if num < 0 {
		return what
	}
	for i := range what {
		if num == 0 {
			return DropString(i, what)
		}
		num--
	}
	// Unreachable
	return ""
}

package gana

// SkipRunes skips the first `num` runes of a string by slicing (underlying array unaffected).
func SkipRunes(num uint, what string) string {
	if uint(len(what)) < num {
		return ""
	}
	for i := range what {
		if num == 0 {
			return SkipString(uint(i), what)
		}
		num--
	}
	// Unreachable
	return ""
}

// DropRunes allocates a new string, with the first `num` runes of a string dropped.
func DropRunes(num uint, what string) string {
	if uint(len(what)) < num {
		return ""
	}
	for i := range what {
		if num == 0 {
			return DropString(uint(i), what)
		}
		num--
	}
	// Unreachable
	return ""
}

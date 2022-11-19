package lazy

// Collect consumes an iterator, returning a slice of all elements consumed.
func Collect[T any](iter Iter[T]) []T {
	result := make([]T, 0)
	for {
		e, ok := iter.Next()
		if !ok {
			break
		}
		result = append(result, e)
	}
	return result
}

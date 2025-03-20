// Package strain provides utility to return from a collection values based on predicate return value.
package strain

// Predicate defines a function that tests an element of type T and returns a boolean.
type Predicate[T any] func(T) bool

// Keep returns a new slice containing elements from items where the predicate returns true.
// If items is empty, it returns an empty slice.
func Keep[T any](items []T, predicate Predicate[T]) []T {
	return filter(items, predicate, true)
}

// Discard returns a new slice containing elements from items where the predicate returns false.
// If items is empty, it returns an empty slice.
func Discard[T any](items []T, predicate Predicate[T]) []T {
	return filter(items, predicate, false)
}

// filter is a helper function that implements the core filtering logic.
// keepMatch determines whether to keep or discard elements when predicate returns true or false.
func filter[T any](items []T, predicate Predicate[T], keepMatch bool) []T {
	result := make([]T, 0, len(items))
	for _, item := range items {
		if predicate(item) == keepMatch {
			result = append(result, item)
		}
	}
	return result
}

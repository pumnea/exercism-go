// Package listops provides utilities for basic list operations.
// Note: Built-in functions like len and append are avoided per exercise requirements.
package listops

// IntList is an abstraction of a list of integers which we can define methods on.
type IntList []int

// Foldl folds (reduces) each item into the accumulator from the left using the given function.
// Example: IntList{1, 2, 3}.Foldl(func(acc, x int) int { return acc + x }, 0) returns 6.
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for i := 0; i < s.Length(); i++ {
		acc = fn(acc, s[i])
	}
	return acc
}

// Foldr folds (reduces) each item into the accumulator from the right using the given function.
// Example: IntList{1, 2, 3}.Foldr(func(x, acc int) int { return x + acc }, 0) returns 6.
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := s.Length() - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

// Filter returns a new list containing only items for which the predicate returns true.
// Example: IntList{1, 2, 3, 4}.Filter(func(x int) bool { return x%2 == 1 }) returns {1, 3}.
func (s IntList) Filter(fn func(int) bool) IntList {
	// Count matches to preallocate result
	count := 0
	for i := 0; i < s.Length(); i++ {
		if fn(s[i]) {
			count++
		}
	}
	// Create result slice
	result := make([]int, count)
	pos := 0
	for i := 0; i < s.Length(); i++ {
		if fn(s[i]) {
			result[pos] = s[i]
			pos++
		}
	}
	return result
}

// Length returns the total number of items in the list.
// Example: IntList{1, 2, 3}.Length() returns 3.
func (s IntList) Length() int {
	length := 0
	for range s {
		length++
	}
	return length
}

// Map returns a new list with the function applied to each item.
// Example: IntList{1, 2, 3}.Map(func(x int) int { return x + 1 }) returns {2, 3, 4}.
func (s IntList) Map(fn func(int) int) IntList {
	result := make([]int, s.Length())
	for i := 0; i < s.Length(); i++ {
		result[i] = fn(s[i])
	}
	return result
}

// Reverse returns a new list with all items in reversed order.
// Example: IntList{1, 2, 3}.Reverse() returns {3, 2, 1}.
func (s IntList) Reverse() IntList {
	result := make([]int, s.Length())
	for i := 0; i < s.Length(); i++ {
		result[s.Length()-1-i] = s[i]
	}
	return result
}

// Append returns a new list with all items from the second list added to the end of the first.
// Example: IntList{1, 2}.Append(IntList{3, 4}) returns {1, 2, 3, 4}.
func (s IntList) Append(lst IntList) IntList {
	result := make([]int, s.Length()+lst.Length())
	for i := 0; i < s.Length(); i++ {
		result[i] = s[i]
	}
	for i := 0; i < lst.Length(); i++ {
		result[s.Length()+i] = lst[i]
	}
	return result
}

// Concat returns a new list combining the receiver with all given lists into one flattened list.
// Example: IntList{1, 2}.Concat([]IntList{{3}, {4, 5}}) returns {1, 2, 3, 4, 5}.
func (s IntList) Concat(lists []IntList) IntList {
	// Calculate total length
	totalLength := s.Length()
	listsCount := 0
	for range lists {
		listsCount++
	}
	for i := 0; i < listsCount; i++ {
		totalLength += lists[i].Length()
	}
	// Create result slice
	result := make([]int, totalLength)
	pos := 0
	// Copy receiver
	for i := 0; i < s.Length(); i++ {
		result[pos] = s[i]
		pos++
	}
	// Copy each list
	for i := 0; i < listsCount; i++ {
		for j := 0; j < lists[i].Length(); j++ {
			result[pos] = lists[i][j]
			pos++
		}
	}
	return result
}

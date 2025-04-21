// Package sublist provides utilities to determine the relationship between two integer slices.
// A sublist is a contiguous subsequence of the other slice.
// For example, Sublist([]int{2, 3}, []int{1, 2, 3}) returns RelationSublist.
package sublist

// Relation type is defined in relations.go.

// Sublist determines if l1 is a sublist, superlist, equal, or unequal to l2.
// Returns RelationEqual, RelationSublist, RelationSuperlist, or RelationUnequal.
func Sublist(l1, l2 []int) Relation {
	if len(l1) == len(l2) && equal(l1, l2) {
		return RelationEqual
	}
	if isSublist(l1, l2) {
		return RelationSublist
	}
	if isSublist(l2, l1) {
		return RelationSuperlist
	}
	return RelationUnequal
}

// equal checks if two slices are identical.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// isSublist checks if small is a contiguous subsequence of large.
func isSublist(small, large []int) bool {
	if len(small) > len(large) {
		return false
	}
	if len(small) == 0 {
		return true // Empty slice is a sublist of any slice
	}
	for i := 0; i <= len(large)-len(small); i++ {
		if equal(small, large[i:i+len(small)]) {
			return true
		}
	}
	return false
}

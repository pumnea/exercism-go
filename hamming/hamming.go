// package hamming provides utility to calculate the Hamming distance in 2 DNA sequances.
package hamming

import "errors"

// Distance return the Hamming distance between 2 DNA sequances or an error if their length differ.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequences length is not equal")
	}
	if len(a) == 0 {
		return 0, nil
	}
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}

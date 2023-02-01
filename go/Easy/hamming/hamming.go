//Package hamming implements functions to calculate the hamming distance between two DNA strands.
package hamming

import "errors"

// Distance calculates and returns the hamming distance between the two strings a and b.
// Both strings should be of the same length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Lengths of strings must be equal")
	}

	hammingDist := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDist += 1
		}
	}

	return hammingDist, nil
}

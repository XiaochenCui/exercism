package hamming

import (
	"errors"
)

// Distance blabla...
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Length of a and b must be equal")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}

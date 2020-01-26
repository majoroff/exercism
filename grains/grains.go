package grains

import (
	"errors"
)

// Square return number of grains on given chessboard square number
// from old puzzle
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("# of square should be in [1, 64]")
	}

	// number of grains on square N is 2^(N-1), let's use it
	power := n - 1
	var grains uint64 = 1 << uint64(power)
	return grains, nil
}

// Total number of grains on chessboard
func Total() uint64 {
	// sum of 2^1 ... 2^n is 2^(n+1)-1
	return 1<<uint64(64) - 1
}

package luhn

import (
	//"fmt"
	"strconv"
	"strings"
)

// Valid : given a number determine whether or not it is valid per the Luhn formula.
func Valid(number string) bool {
	// remove all spaces
	number = strings.ReplaceAll(number, " ", "")
	if len(number) < 2 {
		return false
	}
	runes := []rune(number)
	lenIsEven := len(runes)%2 == 0

	var sum int

	for i := 0; i < len(runes); i++ {
		symbol, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			// spaces removed, only digits are allowed
			return false
		}

		if lenIsEven && i%2 == 0 || !lenIsEven && i%2 != 0 {
			symbol = symbol * 2
			if symbol > 9 {
				symbol = symbol - 9
			}
		}

		sum += symbol
	}
	return sum%10 == 0
}

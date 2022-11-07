package problem_1323

import (
	"math"
	"strconv"
)

func maximum69Number(num int) int {
	result := num
	numAsString := strconv.Itoa(num)

	for i := 0; i < len(numAsString); i++ {
		newNumber := num
		factor := 3 * int(math.Pow10(len(numAsString)-i-1))

		if numAsString[i] == '6' {
			newNumber += factor
		} else {
			newNumber -= factor
		}

		if newNumber > result {
			result = newNumber
		}
	}

	return result
}

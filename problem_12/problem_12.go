package problem_12

import (
	"strconv"
	"strings"
)

func intToRoman(num int) string {
	var result string

	numAsString := strconv.Itoa(num)

	if len(numAsString) > 3 {
		thousand, _ := strconv.Atoi(numAsString[len(numAsString)-4 : len(numAsString)-3])

		result = result + createThousand(thousand*1000)
	}

	if len(numAsString) > 2 {
		hundred, _ := strconv.Atoi(numAsString[len(numAsString)-3 : len(numAsString)-2])

		result = result + createHundred(hundred*100)
	}

	if len(numAsString) > 1 {
		ten, _ := strconv.Atoi(numAsString[len(numAsString)-2 : len(numAsString)-1])

		result = result + createTen(ten*10)
	}

	one, _ := strconv.Atoi(numAsString[len(numAsString)-1:])
	result = result + createOne(one)

	return result
}

func createThousand(num int) string {
	return strings.Repeat("M", num/1000)
}

func createHundred(num int) string {
	switch {
	case num == 900:
		return "CM"
	case num == 500:
		return "D"
	case num == 400:
		return "CD"
	case num > 500 && num < 900:
		return "D" + strings.Repeat("C", (num-500)/100)
	case num < 400:
		return strings.Repeat("C", num/100)
	default:
		return ""
	}
}

func createTen(num int) string {
	switch {
	case num == 90:
		return "XC"
	case num == 50:
		return "L"
	case num == 40:
		return "XL"
	case num > 50 && num < 90:
		return "L" + strings.Repeat("X", (num-50)/10)
	case num < 40:
		return strings.Repeat("X", num/10)
	default:
		return ""
	}
}

func createOne(num int) string {
	switch {
	case num == 9:
		return "IX"
	case num == 5:
		return "V"
	case num == 4:
		return "IV"
	case num > 5 && num < 9:
		return "V" + strings.Repeat("I", num-5)
	case num < 4:
		return strings.Repeat("I", num)
	default:
		return ""
	}
}

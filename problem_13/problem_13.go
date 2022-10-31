package problem_13

func romanToInt(s string) int {
	result := 0

	var romanToNumberMap = map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	for i := 0; i < len(s); i++ {
		if i+2 <= len(s) {
			letterValue, letterExist := romanToNumberMap[s[i:i+2]]
			if letterExist {
				result += letterValue
				i++
			} else {
				result += romanToNumberMap[string(s[i])]
			}
		} else {
			result += romanToNumberMap[string(s[i])]
		}
	}

	return result
}

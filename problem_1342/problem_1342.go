package problem_1342

func numberOfSteps(num int) int {
	steps := 0

	for num != 0 {
		steps++

		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
	}

	return steps
}

func numberOfStepsBinary(num int) int {
	steps := 0

	for num != 0 {
		steps++

		if (num & 1) == 0 {
			num >>= 1
		} else {
			num--
		}
	}

	return steps
}

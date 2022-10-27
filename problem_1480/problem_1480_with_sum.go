package problem_1480

func runningSumWithSum(nums []int) []int {
	var result []int

	sum := 0

	for _, value := range nums {
		sum += value
		result = append(result, sum)
	}

	return result
}

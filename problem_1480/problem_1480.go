package problem_1480

func runningSum(nums []int) []int {
	numsSum := make([]int, len(nums))

	for index, value := range nums {
		if index > 0 {
			numsSum[index] = numsSum[index-1] + value
		} else {
			numsSum[index] = value
		}
	}

	return numsSum
}

func runningSumWithSum(nums []int) []int {
	var result []int

	sum := 0

	for _, value := range nums {
		sum += value
		result = append(result, sum)
	}

	return result
}

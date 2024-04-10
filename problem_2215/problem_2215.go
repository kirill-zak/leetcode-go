package problem2215

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map, nums2Map := make(map[int]bool, len(nums1)), make(map[int]bool, len(nums2))
	num1Result, num2Result := make([]int, 0, len(nums1)), make([]int, 0, len(nums2))

	for _, num := range nums1 {
		nums1Map[num] = true
	}

	for _, num := range nums2 {
		nums2Map[num] = true
	}

	for num1Value := range nums1Map {
		if _, exist := nums2Map[num1Value]; !exist {
			num1Result = append(num1Result, num1Value)
		}
	}

	for num2Value := range nums2Map {
		if _, exist := nums1Map[num2Value]; !exist {
			num2Result = append(num2Result, num2Value)
		}
	}

	return [][]int{num1Result, num2Result}
}

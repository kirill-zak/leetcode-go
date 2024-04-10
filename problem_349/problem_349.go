package problem349

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	resultLen := 0
	if len(nums1) > len(nums2) {
		resultLen = len(nums2)
	} else {
		resultLen = len(nums1)
	}

	result := make([]int, 0, resultLen)

	nums1Map := make(map[int]bool, len(nums1))
	for _, num := range nums1 {
		nums1Map[num] = true
	}

	nums2Map := make(map[int]bool, len(nums2))
	for _, num := range nums2 {
		nums2Map[num] = true
	}

	for num := range nums1Map {
		if _, exist := nums2Map[num]; exist {
			result = append(result, num)
		}
	}

	return result
}

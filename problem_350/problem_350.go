package problem350

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
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

	sort.Ints(nums1)
	sort.Ints(nums2)

	if len(nums1) > len(nums2) {
		nums1Start := 0
		for _, num := range nums2 {
			if nums1Start == len(nums1) {
				break
			}

			for i := nums1Start; i < len(nums1); i++ {
				if num == nums1[i] {
					result = append(result, num)
					nums1Start = i + 1
					break
				}
			}
		}
	} else {
		nums2Start := 0
		for _, num := range nums1 {
			if nums2Start == len(nums2) {
				break
			}

			for i := nums2Start; i < len(nums2); i++ {
				if num == nums2[i] {
					result = append(result, num)
					nums2Start = i + 1
					break
				}
			}
		}
	}

	return result
}

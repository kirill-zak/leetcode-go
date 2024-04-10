package problem26

func removeDuplicates(nums []int) int {
	i := 0
	for {
		if i >= len(nums) {
			break
		}

		if i == 0 {
			i++

			continue
		}

		if nums[i] == nums[i-1] {
			if i+1 >= len(nums) {
				nums = nums[:i]
				break
			}

			nums = append(nums[:i], nums[i+1:]...)

			continue
		}

		i++
	}

	return len(nums)
}

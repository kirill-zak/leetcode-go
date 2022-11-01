package problem_986

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) < 1 || len(secondList) < 1 {
		return [][]int{}
	}

	var result [][]int

	minSecond := 0

	for i := 0; i < len(firstList); i++ {
		for j := minSecond; j < len(secondList); j++ {
			if secondList[j][1] < firstList[i][0] {
				minSecond = j + 1
			} else if firstList[i][1] < secondList[j][0] {
				j = len(secondList)
			} else {
				interval, isIntersect := getIntersection(firstList[i], secondList[j])
				if isIntersect {
					result = append(result, interval)
				}
			}
		}
	}

	return result
}

func getIntersection(first []int, second []int) ([]int, bool) {
	switch {
	case first[0] == second[0] && first[1] == second[1]:
		return first, true
	case first[0] > second[1] || first[1] < second[0]:
		return []int{}, false
	case second[0] >= first[0] && second[1] <= first[1]:
		return second, true
	case second[0] < first[0] && first[1] < second[1]:
		return first, true
	case second[0] < first[0] && second[1] <= first[1]:
		return []int{first[0], second[1]}, true
	case second[0] >= first[0] && second[1] > first[1]:
		return []int{second[0], first[1]}, true
	}

	return []int{}, false
}

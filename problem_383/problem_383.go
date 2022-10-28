package problem_383

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	var magazineMap = make(map[uint8]int)

	for i := 0; i < len(magazine); i++ {
		_, letterExist := magazineMap[magazine[i]]
		if !letterExist {
			magazineMap[magazine[i]] = 0
		}

		magazineMap[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		_, letterExist := magazineMap[ransomNote[i]]
		if !letterExist {
			return false
		}

		magazineMap[ransomNote[i]]--

		if magazineMap[ransomNote[i]] < 1 {
			delete(magazineMap, ransomNote[i])
		}
	}

	return true
}

package problem_14

func longestCommonPrefix(strs []string) string {
	var prefix string
	if len(strs) > 0 {
		prefix = strs[0]
	}

	for _, str := range strs {
		if len(str) < 1 {
			return ""
		}

		startLen := len(prefix)
		if startLen > len(str) {
			prefix = prefix[0:len(str)]
			startLen = len(str)
		}

		for i := startLen; i > 0; i-- {
			if prefix == str[0:i] {
				break
			}

			prefix = prefix[0 : len(prefix)-1]
		}

		if prefix == "" {
			break
		}
	}

	return prefix
}

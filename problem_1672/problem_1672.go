package problem_1672

func maximumWealth(accounts [][]int) int {
	wealth := 0

	for _, account := range accounts {
		accountWealth := 0
		for _, item := range account {
			accountWealth += item
		}

		if accountWealth > wealth {
			wealth = accountWealth
		}
	}

	return wealth
}

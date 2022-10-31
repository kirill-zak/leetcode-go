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

func maximumWealthWithChannel(accounts [][]int) (wealth int) {
	calculateAccountWealth := func(account []int, wealthSum chan int) {
		accountWealth := 0
		for _, item := range account {
			accountWealth += item
		}

		wealthSum <- accountWealth
	}

	for _, account := range accounts {
		wealthSum := make(chan int)

		go calculateAccountWealth(account, wealthSum)

		accountWealth := <-wealthSum
		if accountWealth > wealth {
			wealth = accountWealth
		}

		close(wealthSum)
	}

	return
}

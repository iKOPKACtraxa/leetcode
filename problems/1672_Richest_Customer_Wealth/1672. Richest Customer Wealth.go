package richestcustomerwealth

func maximumWealth(accounts [][]int) int {
	var totalMaximum int
	for _, account := range accounts {
		var maximum int
		for _, sum := range account {
			maximum += sum
		}
		if maximum > totalMaximum {
			totalMaximum = maximum
		}
	}
	return totalMaximum
}

// Time complexity: O(n*m)
// Space complexity: O(1)

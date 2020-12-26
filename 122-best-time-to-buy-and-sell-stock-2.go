package main

/*
# Peak Valley Approach

*/
func maxProfit(prices []int) int {
	var (
		i         int
		maxProfit int
	)
	valley := prices[0]
	peak := prices[0]

	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		maxProfit += peak - valley
	}
	return maxProfit
}

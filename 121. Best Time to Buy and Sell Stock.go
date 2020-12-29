package main

func maxProfit1(prices []int) int {
	var (
		min    int
		profit int
	)
	for i, p := range prices {
		if i == 0 || min > p {
			min = p
		} else if p-min > profit {
			profit = p - min
		}
	}
	return profit
}

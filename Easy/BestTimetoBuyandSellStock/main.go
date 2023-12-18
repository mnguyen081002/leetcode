package main

func maxProfit(prices []int) int {
	maxProfit := 0
	min := prices[0]
	for _, v := range prices {
		if v < min {
			min = v
		}
		max := v - min
		if max > maxProfit {
			maxProfit = max
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit(prices)
}

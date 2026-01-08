package LeetCode

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for _, num := range prices {
		if num < minPrice {
			minPrice = num
		}
		profit := num - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

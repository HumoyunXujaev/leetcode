func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if maxProfit < prices[i]-minPrice {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

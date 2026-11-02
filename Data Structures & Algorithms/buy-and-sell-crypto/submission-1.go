func maxProfit(prices []int) int {
	left,right := 0,0
	maxProfit := 0
	for i, _ := range prices {
		if i == 0 {
			continue
		}
		if prices[i] > prices[right] {
			right = i
			newProfit := prices[right]-prices[left]
			if newProfit > maxProfit {
				maxProfit = newProfit
			}
		}
		if prices[i] < prices[left] {
			left = i
			right = i
		}
	}
	return maxProfit
}

package Best_Time_To_Buy_And_Sell_Stock

max proProfit(prices []int) int{
	lenght := len(prices)
	minPrice := prices[0]
	maxProfit := 0
	for i:=0; i<lenght; i++{
		if prices[i] < minPrice{
			minPrice = prices[i]
		}else{
			profit := prices[i] - minPrice
			if profit > maxProfit{
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

package day1

func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][0]
}

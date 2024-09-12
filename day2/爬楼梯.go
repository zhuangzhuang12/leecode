package day2

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i >= 2 {
			dp[i] = dp[i-2] + 1
		}
	}
	return dp[n]
}

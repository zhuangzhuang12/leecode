package day1

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] >= nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 1
	for _, i := range dp {
		res = max(res, i)
	}
	return res
}

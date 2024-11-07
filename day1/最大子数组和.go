package day1

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], dp[i-1])
	}
	res := -1 << 31
	for _, i := range dp {
		res = max(res, i)
	}

	return res
}
